package expr

import "fmt"

// BetweenExp represents a BETWEEN or a NOT BETWEEN expression.
type BetweenExp struct {
	Col      string
	From, To interface{}
	Not      bool
}

// Build converts an expression into a SQL fragment.
func (e BetweenExp) Build(params Params) string {
	between := "BETWEEN"
	if e.Not {
		between = "NOT BETWEEN"
	}
	name1 := fmt.Sprintf("p%v", len(params))
	name2 := fmt.Sprintf("p%v", len(params)+1)
	params[name1] = e.From
	params[name2] = e.To
	return fmt.Sprintf("%v %v {:%v} AND {:%v}", e.Col, between, name1, name2)
}

// Between generates a BETWEEN expression.
// For example, Between("age", 10, 30) generates: "age" BETWEEN 10 AND 30
func Between(col string, from, to interface{}) Expr {
	return BetweenExp{Col: col, From: from, To: to, Not: false}
}

// NotBetween generates a NOT BETWEEN expression.
// For example, NotBetween("age", 10, 30) generates: "age" NOT BETWEEN 10 AND 30
func NotBetween(col string, from, to interface{}) Expr {
	return BetweenExp{Col: col, From: from, To: to, Not: true}
}
