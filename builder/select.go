package builder

import (
	"github.com/rumis/seal/expr"
	"github.com/rumis/seal/utils"
)

// Select represents a SELECT query
// It can be built into a select sql clauses and params by calling the ToSql method.
type Select struct {
	b            Builder
	selects      []expr.Expr
	distinct     bool
	selectOption string
	from         []string
	where        expr.Expr
	join         []expr.JoinInfo
	orderBy      []string
	groupBy      []string
	having       expr.Expr
	limit        int64
	offset       int64
}

// NewSelect
func NewSelect(b Builder) *Select {
	return &Select{
		b: b,
	}
}

// Select specifies the columns to be selected.
func (s *Select) Select(cols ...string) *Select {
	s.selects = append(s.selects, expr.ColumnExp{
		Columns: cols,
		Table:   "",
	})
	return s
}

// AndSelect adds additional columns to be selected.
func (s *Select) AndSelect(table string, cols ...string) *Select {
	s.selects = append(s.selects, expr.ColumnExp{
		Columns: cols,
		Table:   utils.AliasName(table),
	})
	return s
}

// Agg add aggregate column to be selected
func (s *Select) Agg(fn string, col string, alias string, table ...string) *Select {
	tableName := ""
	if len(table) > 0 {
		tableName = table[0]
	}
	s.selects = append(s.selects, expr.AggExp{
		Column: col,
		Func:   fn,
		Alias:  alias,
		Table:  tableName,
	})
	return s
}

// Distinct specifies whether to select columns distinctively.
// By default, distinct is false.
func (s *Select) Distinct(v bool) *Select {
	s.distinct = v
	return s
}

// SelectOption specifies additional option that should be append to "SELECT".
func (s *Select) SelectOption(option string) *Select {
	s.selectOption = option
	return s
}

// From specifies which tables to select from.
// First table name will combine with selects
func (s *Select) From(tables ...string) *Select {
	s.from = tables
	if len(tables) > 0 && len(s.selects) > 0 {
		if colExp, ok := s.selects[0].(expr.ColumnExp); ok && colExp.Table == "" {
			colExp.Table = utils.AliasName(tables[0])
			s.selects[0] = colExp
		}
	}
	return s
}

// Where specifies the WHERE condition.
// if Where is set before, AndWhere is called
func (s *Select) Where(e expr.Expr) *Select {
	if s.where == nil {
		s.where = e
		return s
	}
	return s.AndWhere(e)
}

// AndWhere concatenates a new WHERE condition with the existing one (if any) using "AND".
func (s *Select) AndWhere(e expr.Expr) *Select {
	s.where = expr.AndOrExp{
		Exps: []expr.Expr{s.where, e},
		Op:   "AND",
	}
	return s
}

// OrWhere concatenates a new WHERE condition with the existing one (if any) using "OR".
func (s *Select) OrWhere(e expr.Expr) *Select {
	s.where = expr.AndOrExp{
		Exps: []expr.Expr{s.where, e},
		Op:   "OR",
	}
	return s
}

// Join specifies a JOIN clause.
// The "typ" parameter specifies the JOIN type (e.g. "INNER JOIN", "LEFT JOIN").
func (s *Select) Join(typ string, table string, on expr.Expr, cols ...string) *Select {
	s.join = append(s.join, expr.JoinInfo{Join: typ, Table: table, On: on})
	s.AndSelect(table, cols...)
	return s
}

// InnerJoin specifies an INNER JOIN clause.
// This is a shortcut method for Join.
func (s *Select) InnerJoin(table string, on expr.Expr, cols ...string) *Select {
	return s.Join("INNER JOIN", table, on, cols...)
}

// LeftJoin specifies a LEFT JOIN clause.
// This is a shortcut method for Join.
func (s *Select) LeftJoin(table string, on expr.Expr, cols ...string) *Select {
	return s.Join("LEFT JOIN", table, on, cols...)
}

// RightJoin specifies a RIGHT JOIN clause.
// This is a shortcut method for Join.
func (s *Select) RightJoin(table string, on expr.Expr, cols ...string) *Select {
	return s.Join("RIGHT JOIN", table, on, cols...)
}

// OrderBy specifies the ORDER BY clause.
// Column names will be properly quoted. A column name can contain "ASC" or "DESC" to indicate its ordering direction.
func (s *Select) OrderBy(cols ...string) *Select {
	s.orderBy = cols
	return s
}

// AndOrderBy appends additional columns to the existing ORDER BY clause.
// Column names will be properly quoted. A column name can contain "ASC" or "DESC" to indicate its ordering direction.
func (s *Select) AndOrderBy(cols ...string) *Select {
	s.orderBy = append(s.orderBy, cols...)
	return s
}

// GroupBy specifies the GROUP BY clause.
// Column names will be properly quoted.
func (s *Select) GroupBy(cols ...string) *Select {
	s.groupBy = cols
	return s
}

// AndGroupBy appends additional columns to the existing GROUP BY clause.
// Column names will be properly quoted.
func (s *Select) AndGroupBy(cols ...string) *Select {
	s.groupBy = append(s.groupBy, cols...)
	return s
}

// Having specifies the HAVING clause.
func (s *Select) Having(e expr.Expr) *Select {
	if s.having == nil {
		s.having = e
		return s
	}
	return s.AndHaving(e)
}

// AndHaving concatenates a new HAVING condition with the existing one (if any) using "AND".
func (s *Select) AndHaving(e expr.Expr) *Select {
	s.having = expr.AndOrExp{
		Exps: []expr.Expr{s.having, e},
		Op:   "AND",
	}
	return s
}

// OrHaving concatenates a new HAVING condition with the existing one (if any) using "OR".
func (s *Select) OrHaving(e expr.Expr) *Select {
	s.having = expr.AndOrExp{
		Exps: []expr.Expr{s.having, e},
		Op:   "OR",
	}
	return s
}

// Limit specifies the LIMIT clause.
// A negative limit means no limit.
func (s *Select) Limit(limit int64) *Select {
	s.limit = limit
	return s
}

// Offset specifies the OFFSET clause.
// A negative offset means no offset.
func (s *Select) Offset(offset int64) *Select {
	s.offset = offset
	return s
}

// build build the sql and params
func (s *Select) build() (string, expr.Params) {
	params := expr.Params{}

	if len(s.from) == 1 {
		// if only one table, remove the table name from column
		for i, colExp := range s.selects {
			switch col := colExp.(type) {
			case expr.ColumnExp:
				col.Table = ""
				s.selects[i] = col
			case expr.AggExp:
				col.Table = ""
				s.selects[i] = col
			}
		}
	}
	clauses := []string{
		s.b.Select(s.selects, s.distinct, s.selectOption),
		s.b.From(s.from),
		s.b.Join(s.join, params),
		s.b.Where(s.where, params),
		s.b.GroupBy(s.groupBy),
		s.b.Having(s.having, params),
		s.b.OrderBy(s.orderBy),
		s.b.Limit(s.limit, s.offset),
	}
	sql := ""
	for _, clause := range clauses {
		if clause != "" {
			if sql == "" {
				sql = clause
			} else {
				sql += " " + clause
			}
		}
	}
	return sql, params
}

// ToSql
func (s *Select) ToSql() (string, []interface{}, error) {
	sql, params := s.build()
	return utils.ReplacePlaceHolders(sql, s.b.Placeholder(), params)
}

// ToExpr build the sql and return an expr
func (s *Select) ToExpr() expr.Expr {
	sql, params := s.build()
	return expr.New(sql, params)
}
