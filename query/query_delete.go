package query

import (
	"github.com/rumis/seal/builder"

	"github.com/rumis/seal/expr"
)

type DeleteQuery struct {
	bd *builder.Delete
	e  Executor
}

func NewDeleteQuery(b builder.Builder, e Executor) *DeleteQuery {
	return &DeleteQuery{
		bd: builder.NewDelete(b),
		e:  e,
	}
}

func (d *DeleteQuery) From(table string) *DeleteQuery {
	d.bd.Table(table)
	return d
}

func (d *DeleteQuery) Where(e expr.Expr) *DeleteQuery {
	d.bd.Where(e)
	return d
}

func (d *DeleteQuery) And(e expr.Expr) *DeleteQuery {
	d.bd.AndWhere(e)
	return d
}

func (d *DeleteQuery) Or(e expr.Expr) *DeleteQuery {
	d.bd.OrWhere(e)
	return d
}

// Exec
func (u *DeleteQuery) Exec(cnt *int64) error {
	sql, args, err := u.bd.ToSql()
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
