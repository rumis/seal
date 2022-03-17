package expr

// Expression represents a DB expression that can be embedded in a SQL statement.
type Expr interface {
	// Build converts an expression into a SQL fragment.
	// If the expression contains binding parameters, they will be added to the given Params.
	Build(Params) string
}

// Params represents a list of parameter values to be bound to a SQL statement.
// The map keys are the parameter names while the map values are the corresponding parameter values.
type Params map[string]interface{}

// JoinInfo contains the specification for a JOIN clause.
type JoinInfo struct {
	Join  string
	Table string
	On    Expr
}

// SelectInfo contains the specification for select columns and table name
type SelectInfo struct {
	Column []string
	Table  string
}

// Exp represents an expression with a SQL fragment and a list of optional binding parameters.
type Exp struct {
	E      string
	Params Params
}

// Build converts an expression into a SQL fragment.
func (e Exp) Build(params Params) string {
	if len(e.Params) == 0 {
		return e.E
	}
	for k, v := range e.Params {
		params[k] = v
	}
	return e.E
}

// New generates an expression with the specified SQL fragment and the optional binding parameters.
func New(e string, params ...Params) Expr {
	if len(params) > 0 {
		return Exp{
			E:      e,
			Params: params[0],
		}
	}
	return Exp{
		E: e,
	}
}
