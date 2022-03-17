package query

import (
	"database/sql"

	"github.com/rumis/seal/utils"
)

// Rows enhances sql.Rows by providing additional data query methods.
// Rows can be obtained by calling Query.Rows(). It is mainly used to populate data row by row.
type Rows struct {
	*sql.Rows
	err error
}

// NewRows generates an Rows instance with *sql.Rows and error which come from db.Query.
func NewRows(rows *sql.Rows, err error) Rows {
	return Rows{rows, err}
}

// AllMap scan all rows and convert to map slice
func (r Rows) AllMap() ([]map[string]interface{}, error) {
	if r.err != nil {
		return nil, r.err
	}
	defer r.Rows.Close()
	cols, err := r.Columns()
	if err != nil {
		return nil, err
	}
	rows := make([]map[string]interface{}, 0)
	for r.Next() {
		var refs []interface{}
		for i := 0; i < len(cols); i++ {
			var t interface{}
			refs = append(refs, &t)
		}
		if err := r.Scan(refs...); err != nil {
			return nil, err
		}
		rowMap := make(map[string]interface{})
		for i, col := range cols {
			rowMap[col] = *refs[i].(*interface{})
		}
		rows = append(rows, rowMap)
	}
	return rows, r.Close()
}

// AllStruct scan all rows and convert to struct slice
func (r Rows) AllStruct(ref interface{}) error {
	rows, err := r.AllMap()
	if err != nil {
		return err
	}
	err = utils.Map2Struct(rows, ref)
	if err != nil {
		return err
	}
	return nil
}

// OneMap scan one row and convert to map
func (r Rows) OneMap() (map[string]interface{}, error) {
	if r.err != nil {
		return nil, r.err
	}
	defer r.Rows.Close()
	cols, err := r.Columns()
	if err != nil {
		return nil, err
	}
	row := make(map[string]interface{})
	if !r.Next() {
		if err := r.Err(); err != nil {
			return row, err
		}
		return row, err
	}
	var refs []interface{}
	for i := 0; i < len(cols); i++ {
		var t interface{}
		refs = append(refs, &t)
	}
	if err := r.Scan(refs...); err != nil {
		return nil, err
	}
	rowMap := make(map[string]interface{})
	for i, col := range cols {
		rowMap[col] = *refs[i].(*interface{})
	}
	return rowMap, r.Close()
}

// OneStruct scan one row and convert to struct
func (r Rows) OneStruct(ref interface{}) error {
	row, err := r.OneMap()
	if err != nil {
		return err
	}
	err = utils.Map2Struct(row, ref)
	if err != nil {
		return err
	}
	return nil
}
