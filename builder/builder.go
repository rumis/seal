package builder

import (
	"github.com/rumis/seal/expr"
)

// Builder supports building SQL statements in a DB-agnostic way.
// Builder mainly provides two sets of query building methods: those building SELECT statements
// and those manipulating DB data or schema (e.g. INSERT statements).
type Builder interface {
	// BuildSelect generates a SELECT clause from the given selected column names.
	Select(cols []expr.SelectInfo, distinct bool, option string) string
	// Insert
	Insert(table string, cols []string, vals [][]interface{}, params expr.Params) string
	// Update
	Update(table string, values map[string]interface{}, params expr.Params) string
	// Delete
	Delete(table string) string
	// From generates a FROM clause from the given tables.
	From(tables []string) string
	// GroupBy generates a GROUP BY clause from the given group-by columns.
	GroupBy(cols []string) string
	// Join generates a JOIN clause from the given join information.
	Join([]expr.JoinInfo, expr.Params) string
	// Where generates a WHERE clause from the given expression.
	Where(expr.Expr, expr.Params) string
	// Having generates a HAVING clause from the given expression.
	Having(expr.Expr, expr.Params) string
	// OrderBy generates the ORDER BY and LIMIT clauses.
	OrderBy([]string) string
	// Limit generates the OFFSET and LIMIT clauses.
	Limit(int64, int64) string
	// Placeholder generates the placeholder char
	Placeholder() string
}
