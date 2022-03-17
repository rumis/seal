package query

import (
	"context"
	"database/sql"
)

// Executor prepares, executes, or queries a SQL statement.
type Executor interface {
	// Exec executes a SQL statement
	Exec(query string, args ...interface{}) (sql.Result, error)
	// ExecContext executes a SQL statement with the given context
	ExecContext(ctx context.Context, query string, args ...interface{}) (sql.Result, error)
	// Query queries a SQL statement
	Query(query string, args ...interface{}) (*sql.Rows, error)
	// QueryContext queries a SQL statement with the given context
	QueryContext(ctx context.Context, query string, args ...interface{}) (*sql.Rows, error)
}
