package builder

import (
	"bytes"
	"fmt"
	"regexp"
	"strings"

	"github.com/rumis/seal/expr"
)

// BuilderStandard stand sql builder
type BuilderStandard struct {
}

var _ Builder = &BuilderStandard{}

// NewStandardBuilder constructure of standardbuilder
func NewStandardBuilder() Builder {
	return BuilderStandard{}
}

// Select generates a SELECT clause from the given selected column names.
func (q BuilderStandard) Select(cols []expr.Expr, distinct bool, option string) string {
	var s bytes.Buffer
	s.WriteString("SELECT ")
	if distinct {
		s.WriteString("DISTINCT ")
	}
	if option != "" {
		s.WriteString(option)
		s.WriteString(" ")
	}
	if len(cols) == 0 {
		s.WriteString("*")
		return s.String()
	}
	column := make([]string, 0, len(cols))
	for _, colExp := range cols {
		column = append(column, colExp.Build(nil))
	}
	s.WriteString(strings.Join(column, ","))
	return s.String()
}

// From generates a FROM clause from the given tables.
func (q BuilderStandard) From(tables []string) string {
	if len(tables) == 0 {
		return ""
	}
	return "FROM " + strings.Join(tables, ", ")
}

// Join generates a JOIN clause from the given join information.
func (q BuilderStandard) Join(joins []expr.JoinInfo, params expr.Params) string {
	if len(joins) == 0 {
		return ""
	}
	parts := []string{}
	for _, join := range joins {
		sql := join.Join + " " + join.Table
		on := ""
		if join.On != nil {
			on = join.On.Build(params)
		}
		if on != "" {
			sql += " ON " + on
		}
		parts = append(parts, sql)
	}
	return strings.Join(parts, " ")
}

// Where generates a WHERE clause from the given expression.
func (q BuilderStandard) Where(e expr.Expr, params expr.Params) string {
	if e != nil {
		if c := e.Build(params); c != "" {
			return "WHERE " + c
		}
	}
	return ""
}

// Having generates a HAVING clause from the given expression.
func (q BuilderStandard) Having(e expr.Expr, params expr.Params) string {
	if e != nil {
		if c := e.Build(params); c != "" {
			return "HAVING " + c
		}
	}
	return ""
}

// GroupBy generates a GROUP BY clause from the given group-by columns.
func (q BuilderStandard) GroupBy(cols []string) string {
	if len(cols) == 0 {
		return ""
	}
	s := ""
	for i, col := range cols {
		if i == 0 {
			s = col
		} else {
			s += ", " + col
		}
	}
	return "GROUP BY " + s
}

var orderRegex = regexp.MustCompile(`\s+((?i)ASC|DESC)$`)

// OrderBy generates the ORDER BY clause.
func (q BuilderStandard) OrderBy(cols []string) string {
	if len(cols) == 0 {
		return ""
	}
	s := ""
	for i, col := range cols {
		if i > 0 {
			s += ", "
		}
		matches := orderRegex.FindStringSubmatch(col)
		if len(matches) == 0 {
			s += col
		} else {
			col := col[:len(col)-len(matches[0])]
			dir := matches[1]
			s += col + " " + dir
		}
	}
	return "ORDER BY " + s
}

// Limit generates the LIMIT clause.
func (q BuilderStandard) Limit(limit int64, offset int64) string {
	if limit < 0 && offset > 0 {
		// most DBMS requires LIMIT when OFFSET is present
		limit = 9223372036854775807 // 2^63 - 1
	}

	sql := ""
	if limit > 0 {
		sql = fmt.Sprintf("LIMIT %v", limit)
	}
	if offset <= 0 {
		return sql
	}
	if sql != "" {
		sql += " "
	}
	return sql + fmt.Sprintf("OFFSET %v", offset)
}

// Delete  generates the DELETE clause.
func (q BuilderStandard) Delete(table string) string {
	sql := "DELETE FROM " + table
	return sql
}

// Update generate the UPDATE clause
func (q BuilderStandard) Update(table string, values map[string]interface{}, params expr.Params) string {
	lines := make([]string, 0, len(values))
	for k, v := range values {
		if e, ok := v.(expr.Expr); ok {
			lines = append(lines, k+"="+e.Build(params))
		} else {
			lines = append(lines, fmt.Sprintf("%v={:p%v}", k, len(params)))
			params[fmt.Sprintf("p%v", len(params))] = v
		}
	}
	return fmt.Sprintf("UPDATE %v SET %v", table, strings.Join(lines, ", "))
}

// Insert generate the insert clause
func (q BuilderStandard) Insert(table string, cols []string, vals [][]interface{}, params expr.Params) string {
	valuesStrings := make([]string, len(vals))
	for r, row := range vals {
		valueStrings := make([]string, len(row))
		for v, val := range row {
			valueStrings[v] = fmt.Sprintf("{:p%v}", len(params))
			params[fmt.Sprintf("p%v", len(params))] = val
		}
		valuesStrings[r] = fmt.Sprintf("(%s)", strings.Join(valueStrings, ","))
	}
	sql := fmt.Sprintf("INSERT INTO %v (%v) VALUES %v",
		table,
		strings.Join(cols, ", "),
		strings.Join(valuesStrings, ", "),
	)
	return sql
}

// Placeholder generates an anonymous parameter placeholder with the given parameter ID.
func (q BuilderStandard) Placeholder() string {
	return "?"
}
