package builder

import (
	"errors"
	"reflect"

	"github.com/rumis/seal/expr"
	"github.com/rumis/seal/options"
	"github.com/rumis/seal/utils"
)

// Insert represents a INSERT query
// It can be built into a insert sql clauses and params by calling the ToSql method.
type Insert struct {
	b  Builder
	eh options.EncodeHookFunc

	cols  []string
	table string
	vals  [][]interface{}
}

// NewInsert
func NewInsert(b Builder, encodeHook options.EncodeHookFunc) *Insert {
	return &Insert{
		b:  b,
		eh: encodeHook,
	}
}

// Into specifies which tables to insert into.
func (i *Insert) Into(table string) *Insert {
	i.table = table
	return i
}

// Columns specifies which columns to insert.
func (i *Insert) Columns(columns ...string) *Insert {
	i.cols = columns
	return i
}

// Values specifies the insert values
// type of vals can be []map[string]interface{}, []struct , [][]interface{}
// if type of vals is [][]interface{}, cols must be set first and be matched
func (i *Insert) Values(vals interface{}) *Insert {
	valsMapFunc := func(val []map[string]interface{}) {
		if len(val) == 0 {
			return
		}
		i.vals = make([][]interface{}, 0, len(val))
		// set columns only when Columns func is not called
		if len(i.cols) == 0 {
			i.cols = make([]string, 0, len(val[0]))
			for k := range val[0] {
				i.cols = append(i.cols, k)
			}
		}
		for _, vm := range val {
			rowVal := make([]interface{}, 0, len(vm))
			for _, k := range i.cols {
				if i.eh != nil {
					encodeRes, _ := i.eh(reflect.TypeOf(vm[k]), vm[k])
					rowVal = append(rowVal, encodeRes)
				} else {
					rowVal = append(rowVal, vm[k])
				}
			}
			i.vals = append(i.vals, rowVal)
		}
	}
	switch val := vals.(type) {
	case []map[string]interface{}:
		valsMapFunc(val)
		return i
	case [][]interface{}:
		i.vals = val
		return i
	}
	v, err := utils.Struct2MapSlice(vals)
	if err != nil {
		return i
	}
	valsMapFunc(v)
	return i
}

// Value specifies the insert value
// type of vals can be map[string]interface{} , struct , []interface{}
// if type of vals is []interface{}, cols must be set first and be matched
func (i *Insert) Value(val interface{}) *Insert {
	valMapFunc := func(val map[string]interface{}) {
		i.cols = make([]string, 0, len(val))
		rowVal := make([]interface{}, 0, len(val))
		for k, v := range val {
			rowVal = append(rowVal, v)
			i.cols = append(i.cols, k)
		}
		i.vals = append(i.vals, rowVal)
	}
	switch v := val.(type) {
	case map[string]interface{}:
		valMapFunc(v)
		return i
	case []interface{}:
		i.vals = append(i.vals, v)
		return i
	}
	vm, err := utils.Struct2Map(val)
	if err != nil {
		return i
	}
	valMapFunc(vm)
	return i
}

// ToSql build the sql clauses and params
func (i *Insert) ToSql() (string, []interface{}, error) {
	if len(i.cols) == 0 {
		return "", nil, errors.New("insert columns not set")
	}
	if len(i.vals) == 0 {
		return "", nil, errors.New("insert value not set")
	}
	params := expr.Params{}
	sql := i.b.Insert(i.table, i.cols, i.vals, params)
	return utils.ReplacePlaceHolders(sql, i.b.Placeholder(), params)
}
