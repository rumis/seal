package expr

// NotExp represents an expression that should prefix "NOT" to a specified expression.
type NotExp struct {
	E Expr
}

// Build converts an expression into a SQL fragment.
func (e NotExp) Build(params Params) string {
	if sql := e.E.Build(params); sql != "" {
		return "NOT (" + sql + ")"
	}
	return ""
}

// Not generates a NOT expression which prefixes "NOT" to the specified expression.
func Not(e Expr) Expr {
	return NotExp{E: e}
}
