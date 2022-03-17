package expr

import "fmt"

// StandardExp represents a noraml expressions
// eg. =,>,>=,<,<=,....
type StandardExp struct {
	Col   string
	Op    string
	Value interface{}
}

// Build converts an expression into a SQL fragment.
func (e StandardExp) Build(params Params) string {
	p1 := fmt.Sprintf("p%v", len(params))
	params[p1] = e.Value
	return fmt.Sprintf("%v%v{:%v}", e.Col, e.Op, p1)
}

// Op generates a Standard expression
// eg. =,>,>=,<,<=,....
func Op(col string, op string, val interface{}) Expr {
	return StandardExp{
		Col:   col,
		Op:    op,
		Value: val,
	}
}
