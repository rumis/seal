package query

import (
	"context"
	"time"

	"github.com/rumis/seal/builder"
	"github.com/rumis/seal/expr"
)

// SelectQuery represents the sql builder of select and base query
type SelectQuery struct {
	bs    *builder.Select
	baseQ Query
}

// NewSelectQuery constructure of SelectQuery
func NewSelectQuery(b builder.Builder, q Query) *SelectQuery {
	return &SelectQuery{
		bs:    builder.NewSelect(b),
		baseQ: q,
	}
}

// Select specifies the columns to be selected.
func (s *SelectQuery) Select(cols ...string) *SelectQuery {
	s.bs.Select(cols...)
	return s
}

// Agg specifies the aggregate function.
func (s *SelectQuery) Agg(fn string, col string, alias string, table ...string) *SelectQuery {
	s.bs.Agg(fn, col, alias, table...)
	return s
}

// Distinct specifies whether to select columns distinctively.
// By default, distinct is false.
func (s *SelectQuery) Distinct(v bool) *SelectQuery {
	s.bs.Distinct(v)
	return s
}

// From specifies which tables to select from.
func (s *SelectQuery) From(table string) *SelectQuery {
	s.bs.From(table)
	return s
}

// Where specifies the WHERE condition.
func (s *SelectQuery) Where(e expr.Expr) *SelectQuery {
	s.bs.Where(e)
	return s
}

// And concatenates a new WHERE condition with the existing one (if any) using "AND".
func (s *SelectQuery) And(e expr.Expr) *SelectQuery {
	s.bs.AndWhere(e)
	return s
}

// Or concatenates a new WHERE condition with the existing one (if any) using "OR".
func (s *SelectQuery) Or(e expr.Expr) *SelectQuery {
	s.bs.OrWhere(e)
	return s
}

// Having specifies the HAVING condition.
func (s *SelectQuery) Having(e expr.Expr) *SelectQuery {
	s.bs.Having(e)
	return s
}

// AndHaving concatenates a new HAVING condition with the existing one (if any) using "AND".
func (s *SelectQuery) AndHaving(e expr.Expr) *SelectQuery {
	s.bs.AndHaving(e)
	return s
}

// OrHaving concatenates a new HAVING expr with the existing one (if any) using "OR".
func (s *SelectQuery) OrHaving(e expr.Expr) *SelectQuery {
	s.bs.OrHaving(e)
	return s
}

// OrderBy specifies the ORDER BY clause.
func (s *SelectQuery) OrderBy(cols ...string) *SelectQuery {
	s.bs.OrderBy(cols...)
	return s
}

// GroupBy specifies the GROUP BY clause.
func (s *SelectQuery) GroupBy(cols ...string) *SelectQuery {
	s.bs.GroupBy(cols...)
	return s
}

// Limit specifies the LIMIT clause.
func (s *SelectQuery) Limit(limit int64) *SelectQuery {
	s.bs.Limit(limit)
	return s
}

// Offset specifies the OFFSET clause.
func (s *SelectQuery) Offset(offset int64) *SelectQuery {
	s.bs.Offset(offset)
	return s
}

// Query queries a SQL statement
func (s *SelectQuery) Query() Rows {

	sTime := time.Now()

	sql, args, err := s.bs.ToSql()

	if s.baseQ.opts.BuildLog != nil {
		s.baseQ.opts.BuildLog(context.Background(), time.Since(sTime), sql, args, err)
	}

	if err != nil {
		return NewRows(nil, err)
	}
	return s.baseQ.Query(sql, args...)
}

// ToExpr return the complete sql string. used for sub sql stmt
func (s *SelectQuery) ToExpr() expr.Expr {
	return s.bs.ToExpr()
}
