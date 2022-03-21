package query

import (
	"context"
	"time"

	"github.com/rumis/seal/builder"

	"github.com/rumis/seal/expr"
)

// DeleteQuery represents the sql builder of delete and base query
type DeleteQuery struct {
	bd    *builder.Delete
	baseQ Query
}

// NewDeleteQuery constructure of DeleteQuery
func NewDeleteQuery(b builder.Builder, q Query) *DeleteQuery {
	return &DeleteQuery{
		bd:    builder.NewDelete(b),
		baseQ: q,
	}
}

// From set table name
func (d *DeleteQuery) From(table string) *DeleteQuery {
	d.bd.Table(table)
	return d
}

// Where generates a WHERE clause from the given expression.
func (d *DeleteQuery) Where(e expr.Expr) *DeleteQuery {
	d.bd.Where(e)
	return d
}

// And concatenates a new WHERE condition with the existing one (if any) using "AND".
func (d *DeleteQuery) And(e expr.Expr) *DeleteQuery {
	d.bd.AndWhere(e)
	return d
}

// Or concatenates a new WHERE condition with the existing one (if any) using "OR".
func (d *DeleteQuery) Or(e expr.Expr) *DeleteQuery {
	d.bd.OrWhere(e)
	return d
}

// Exec executes a SQL statement
func (u *DeleteQuery) Exec(cnt *int64) error {
	sTime := time.Now()

	sql, args, err := u.bd.ToSql()

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
