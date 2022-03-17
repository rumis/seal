package expr

// ExistsExp represents an EXISTS or NOT EXISTS expression.
type ExistsExp struct {
	Exp Expr
	Not bool
}

// Build converts an expression into a SQL fragment.
func (e ExistsExp) Build(params Params) string {
	sql := e.Exp.Build(params)
	if sql == "" {
		if e.Not {
			return ""
		}
		return "0=1"
	}
	if e.Not {
		return "NOT EXISTS (" + sql + ")"
	}
	return "EXISTS (" + sql + ")"
}

// Exists generates an EXISTS expression by prefixing "EXISTS" to the given expression.
func Exists(exp Expr) Expr {
	return ExistsExp{Exp: exp, Not: false}
}

// NotExists generates an EXISTS expression by prefixing "NOT EXISTS" to the given expression.
func NotExists(exp Expr) Expr {
	return ExistsExp{Exp: exp, Not: true}
}
