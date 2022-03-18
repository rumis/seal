package query

import (
	"github.com/rumis/seal/builder"
	"github.com/rumis/seal/expr"
)

type SelectQuery struct {
	bs *builder.Select
	e  Executor
}

// NewSelectQuery
func NewSelectQuery(b builder.Builder, e Executor) *SelectQuery {
	return &SelectQuery{
		bs: builder.NewSelect(b),
		e:  e,
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

// Query
func (s *SelectQuery) Query() Rows {
	sql, args, err := s.bs.ToSql()
	if err != nil {
		return NewRows(nil, err)
	}
	rows, err := s.e.Query(sql, args...)
	return NewRows(rows, err)
}

// ToExpr return the complete sql string. used for sub sql stmt
func (s *SelectQuery) ToExpr() expr.Expr {
	return s.bs.ToExpr()
}
