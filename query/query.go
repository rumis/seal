package query

import (
	"context"
	"database/sql"
	"time"

	"github.com/rumis/seal/builder"
	"github.com/rumis/seal/expr"
	"github.com/rumis/seal/options"
)

// Query represents the sql builder and exector. it is the parent class of db and tx
type Query struct {
	b    builder.Builder
	e    Executor
	opts *options.SealOptions
}

// NewQuery generate a base query instance
func NewQuery(b builder.Builder, e Executor, opts *options.SealOptions) Query {
	return Query{b, e, opts}
}

// Builder return the builder
func (q Query) Builder() builder.Builder {
	return q.b
}

// Options return the sealoptions
func (q Query) Options() *options.SealOptions {
	return q.opts
}

// Insert generate the insert query
func (q Query) Insert(table string) *InsertQuery {
	return NewInsertQuery(q.b, q).Into(table)
}

// Update generate the update query
func (q Query) Update(table string) *UpdateQuery {
	return NewUpdateQuery(q.b, q).Table(table)
}

// Delete generate the delete query
func (q Query) Delete(table string) *DeleteQuery {
	return NewDeleteQuery(q.b, q).From(table)
}

// Select generate the select query
func (q Query) Select(cols ...string) *SelectQuery {
	return NewSelectQuery(q.b, q).Select(cols...)
}

// SubQuery generate a sub select sql expr
func (q Query) SubQuery(sub func(q *SelectQuery)) expr.Expr {
	sq := NewSelectQuery(q.b, q)
	sub(sq)
	return sq.ToExpr()
}

// Select generate the select query
func (q Query) Count(col string) *SelectQuery {
	return NewSelectQuery(q.b, q).Agg("COUNT", col, ALIAS_AGG_COUNT)
}

// ExecContext exec a raw sql with context
func (q Query) ExecContext(ctx context.Context, sql string, args ...interface{}) sql.Result {
	sTime := time.Now()
	result, err := q.e.ExecContext(ctx, sql, args...)

	if q.opts.ExecLog != nil {
		q.opts.ExecLog(ctx, time.Since(sTime), sql, args, err)
	}

	return NewExecResult(result, err)
}

// QueryContext exec a raw query sql with context
func (q Query) QueryContext(ctx context.Context, sql string, args ...interface{}) Rows {
	sTime := time.Now()

	rows, err := q.e.QueryContext(ctx, sql, args...)

	if q.opts.ExecLog != nil {
		q.opts.ExecLog(ctx, time.Since(sTime), sql, args, err)
	}

	if err != nil {
		return NewRows(nil, err)
	}
	return NewRows(rows, err)
}
