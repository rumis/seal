package expr

import "strings"

// AndOrExp represents an expression that concatenates multiple expressions using either "AND" or "OR".
type AndOrExp struct {
	Exps []Expr
	Op   string
}

// Build converts an expression into a SQL fragment.
func (e AndOrExp) Build(params Params) string {
	if len(e.Exps) == 0 {
		return ""
	}
	var parts []string
	for _, a := range e.Exps {
		if a == nil {
			continue
		}
		if sql := a.Build(params); sql != "" {
			parts = append(parts, sql)
		}
	}
	if len(parts) == 1 {
		return parts[0]
	}
	return strings.Join(parts, " "+e.Op+" ")
}

// And generates an AND expression which concatenates the given expressions with "AND".
func And(exps ...Expr) Expr {
	return AndOrExp{Exps: exps, Op: "AND"}
}

// Or generates an OR expression which concatenates the given expressions with "OR".
func Or(exps ...Expr) Expr {
	return AndOrExp{Exps: exps, Op: "OR"}
}
