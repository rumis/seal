package query

import (
	"context"
	"time"

	"github.com/rumis/seal/builder"
)

// InsertQuery represents the sql builder of insert and base query
type InsertQuery struct {
	bi    *builder.Insert
	baseQ Query
}

// NewInsertQuery constructure of insertquery
func NewInsertQuery(b builder.Builder, q Query) *InsertQuery {
	return &InsertQuery{
		bi:    builder.NewInsert(b, q.opts.EncodeHook),
		baseQ: q,
	}
}

// Into set the table
func (i *InsertQuery) Into(table string) *InsertQuery {
	i.bi.Into(table)
	return i
}

// Columns set columns name when the data type is []interface{} or [][]interface{}
func (i *InsertQuery) Columns(columns ...string) *InsertQuery {
	i.bi.Columns(columns...)
	return i
}

// Values set the data will insert into the table
func (i *InsertQuery) Values(vals interface{}) *InsertQuery {
	i.bi.Values(vals)
	return i
}

// Value set the data will insert into the table
func (i *InsertQuery) Value(val interface{}) *InsertQuery {
	i.bi.Value(val)
	return i
}

// Exec executes a SQL statement
func (u *InsertQuery) Exec(ctx context.Context, lastId *int64) error {
	sTime := time.Now()

	sql, args, err := u.bi.ToSql()

	if u.baseQ.opts.BuildLog != nil {
		u.baseQ.opts.BuildLog(ctx, time.Since(sTime), sql, args, err)
	}

	if err != nil {
		return err
	}
	result := u.baseQ.ExecContext(ctx, sql, args...)
	*lastId, err = result.LastInsertId()
	if err != nil {
		return err
	}
	return nil
}
