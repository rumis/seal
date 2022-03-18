package expr

import (
	"strings"
)

// ColumnExp represents an expression that for columns.
type ColumnExp struct {
	Columns []string
	Table   string
}

// Build converts an expression into a SQL fragment.
func (e ColumnExp) Build(params Params) string {
	column := make([]string, 0, len(e.Columns))
	for _, col := range e.Columns {
		if e.Table != "" {
			col = e.Table + "." + col
		}
		column = append(column, col)
	}
	return strings.Join(column, ",")
}
