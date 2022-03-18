package query

import (
	"context"
	"database/sql"

	"github.com/rumis/seal/builder"
	"github.com/rumis/seal/expr"
)

// Query represents the sql builder and exector. it is the parent class of db and tx
type Query struct {
	b builder.Builder
	e Executor
}

// NewQuery generate a base query instance
func NewQuery(b builder.Builder, e Executor) Query {
	return Query{b, e}
}

// Builder return the builder
func (q Query) Builder() builder.Builder {
	return q.b
}

// Insert generate the insert query
func (q Query) Insert(table string) *InsertQuery {
	return NewInsertQuery(q.b, q.e).Into(table)
}

// Update generate the update query
func (q Query) Update(table string) *UpdateQuery {
	return NewUpdateQuery(q.b, q.e).Table(table)
}

// Delete generate the delete query
func (q Query) Delete(table string) *DeleteQuery {
	return NewDeleteQuery(q.b, q.e).From(table)
}

// Select generate the select query
func (q Query) Select(cols ...string) *SelectQuery {
	return NewSelectQuery(q.b, q.e).Select(cols...)
}

// SubQuery generate a sub select sql expr
func (q Query) SubQuery(sub func(q *SelectQuery)) expr.Expr {
	sq := NewSelectQuery(q.b, q.e)
	sub(sq)
	return sq.ToExpr()
}

// Select generate the select query
func (q Query) Count(col string) *SelectQuery {
	return NewSelectQuery(q.b, q.e).Agg("COUNT", col, ALIAS_AGG_COUNT)
}

// Exec exec a raw sql
func (q Query) Exec(sql string, args ...interface{}) sql.Result {
	result, err := q.e.Exec(sql, args...)
	return NewExecResult(result, err)
}

// ExecContext exec a raw sql with context
func (q Query) ExecContext(ctx context.Context, query string, args ...interface{}) sql.Result {
	result, err := q.e.ExecContext(ctx, query, args...)
	return NewExecResult(result, err)
}

// Query exec a raw query sql
func (q Query) Query(sql string, args ...interface{}) Rows {
	rows, err := q.e.Query(sql, args...)
	if err != nil {
		return NewRows(nil, err)
	}
	return NewRows(rows, err)
}

// QueryContext exec a raw query sql with context
func (q Query) QueryContext(ctx context.Context, sql string, args ...interface{}) Rows {
	rows, err := q.e.QueryContext(ctx, sql, args...)
	if err != nil {
		return NewRows(nil, err)
	}
	return NewRows(rows, err)
}
