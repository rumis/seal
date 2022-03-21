package query

import (
	"context"
	"time"

	"github.com/rumis/seal/builder"
	"github.com/rumis/seal/expr"
)

// UpdateQuery represents the sql builder of update and base query
type UpdateQuery struct {
	bu    *builder.Update
	baseQ Query
}

// NewUpdateQuery constructure of UpdateQuery
func NewUpdateQuery(b builder.Builder, q Query) *UpdateQuery {
	return &UpdateQuery{
		bu:    builder.NewUpdate(b),
		baseQ: q,
	}
}

// Table set table name
func (u *UpdateQuery) Table(table string) *UpdateQuery {
	u.bu.Table(table)
	return u
}

// Value set the data which will be update
func (u *UpdateQuery) Value(val interface{}) *UpdateQuery {
	u.bu.Value(val)
	return u
}

// Where  generates a WHERE clause from the given expression.
func (u *UpdateQuery) Where(e expr.Expr) *UpdateQuery {
	u.bu.Where(e)
	return u
}

// And concatenates a new WHERE condition with the existing one (if any) using "AND".
func (u *UpdateQuery) And(e expr.Expr) *UpdateQuery {
	u.bu.AndWhere(e)
	return u
}

// Or concatenates a new WHERE condition with the existing one (if any) using "OR".
func (u *UpdateQuery) Or(e expr.Expr) *UpdateQuery {
	u.bu.OrWhere(e)
	return u
}

// Exec executes a SQL statement
func (u *UpdateQuery) Exec(cnt *int64) error {

	sTime := time.Now()

	sql, args, err := u.bu.ToSql()

	if u.baseQ.opts.BuildLog != nil {
		u.baseQ.opts.BuildLog(context.Background(), time.Since(sTime), sql, args, err)
	}

	if err != nil {
		return err
	}
	result := u.baseQ.Exec(sql, args...)
	*cnt, err = result.RowsAffected()
	if err != nil {
		return err
	}
	return nil
}
