package query

import (
	"github.com/rumis/seal/builder"
)

// InsertQuery
type InsertQuery struct {
	bi *builder.Insert
	e  Executor
}

// NewUpdateQuery
func NewInsertQuery(b builder.Builder, e Executor) *InsertQuery {
	return &InsertQuery{
		bi: builder.NewInsert(b),
		e:  e,
	}
}

// Into
func (i *InsertQuery) Into(table string) *InsertQuery {
	i.bi.Into(table)
	return i
}

// Columns
func (i *InsertQuery) Columns(columns ...string) *InsertQuery {
	i.bi.Columns(columns...)
	return i
}

// Values
func (i *InsertQuery) Values(vals interface{}) *InsertQuery {
	i.bi.Values(vals)
	return i
}

// Value
func (i *InsertQuery) Value(val interface{}) *InsertQuery {
	i.bi.Value(val)
	return i
}

// Exec
func (u *InsertQuery) Exec(cnt *int64) error {
	sql, args, err := u.bi.ToSql()
	if err != nil {
		return err
	}
	result, err := u.e.Exec(sql, args...)
	if err != nil {
		return err
	}
	*cnt, err = result.LastInsertId()
	if err != nil {
		return err
	}
	return nil
}
