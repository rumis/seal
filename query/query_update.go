package query

import (
	"github.com/rumis/seal/builder"
	"github.com/rumis/seal/expr"
)

// UpdateQuery
type UpdateQuery struct {
	bu *builder.Update
	e  Executor
}

// NewUpdateQuery
func NewUpdateQuery(b builder.Builder, e Executor) *UpdateQuery {
	return &UpdateQuery{
		bu: builder.NewUpdate(b),
		e:  e,
	}
}

// Table
func (u *UpdateQuery) Table(table string) *UpdateQuery {
	u.bu.Table(table)
	return u
}

// Value
func (u *UpdateQuery) Value(val interface{}) *UpdateQuery {
	u.bu.Value(val)
	return u
}

// Where
func (u *UpdateQuery) Where(e expr.Expr) *UpdateQuery {
	u.bu.Where(e)
	return u
}

// Or
func (u *UpdateQuery) And(e expr.Expr) *UpdateQuery {
	u.bu.AndWhere(e)
	return u
}

// Or
func (u *UpdateQuery) Or(e expr.Expr) *UpdateQuery {
	u.bu.OrWhere(e)
	return u
}

// Exec
func (u *UpdateQuery) Exec(cnt *int64) error {
	sql, args, err := u.bu.ToSql()
	if err != nil {
		return err
	}
	result, err := u.e.Exec(sql, args...)
	if err != nil {
		return err
	}
	*cnt, err = result.RowsAffected()
	if err != nil {
		return err
	}
	return nil
}
