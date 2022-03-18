package expr

import (
	"fmt"

	"github.com/rumis/seal/utils"
)

// AggExp represents an expression that for Aggregate function clause.
type AggExp struct {
	Column string
	Alias  string
	Func   string
	Table  string
}

// Build converts an expression into a SQL fragment.
func (e AggExp) Build(params Params) string {
	col := e.Column
	if e.Table != "" {
		col = utils.AliasName(e.Table) + "." + col
	}
	return fmt.Sprintf("%v(%v) AS %v", e.Func, col, e.Alias)
}

// Aggregate generates a aggregate function expression.
// count,sum,average,max,min
func Aggregate(fn string, col string, alias ...string) Expr {
	var a string
	if len(alias) > 0 {
		a = alias[0]
	}
	var t string
	if len(alias) > 1 {
		t = alias[1]
	}
	return AggExp{
		Column: col,
		Func:   fn,
		Alias:  a,
		Table:  t,
	}
}
