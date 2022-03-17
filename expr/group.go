package expr

import "strings"

// GroupExp represents an expression that concatenates multiple expressions using either "AND" or "OR".
type GroupExp struct {
	Exps []Expr
	Op   string
}

// Build converts an expression into a SQL fragment.
func (e GroupExp) Build(params Params) string {
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
	return "(" + strings.Join(parts, " "+e.Op+" ") + ")"
}

// Group generates multi expression with parentheses
func Group(op string, exps ...Expr) Expr {
	return GroupExp{
		Op:   op,
		Exps: exps,
	}
}
