package expr

import (
	"fmt"
	"strings"
)

// InExp represents an "IN" or "NOT IN" expression.
type InExp struct {
	Col    string
	Values []interface{}
	Not    bool
}

// Build converts an expression into a SQL fragment.
func (e InExp) Build(params Params) string {
	if len(e.Values) == 0 {
		if e.Not {
			return ""
		}
		return "0=1"
	}
	var values []string
	for _, value := range e.Values {
		switch v := value.(type) {
		case nil:
			values = append(values, "NULL")
		case Expr:
			sql := v.Build(params)
			values = append(values, sql)
		default:
			name := fmt.Sprintf("p%v", len(params))
			params[name] = value
			values = append(values, "{:"+name+"}")
		}
	}
	// if only one value, operate IN fall back to Equal
	if len(values) == 1 {
		if e.Not {
			return e.Col + "<>" + values[0]
		}
		return e.Col + "=" + values[0]
	}
	in := "IN"
	if e.Not {
		in = "NOT IN"
	}
	return fmt.Sprintf("%v %v (%v)", e.Col, in, strings.Join(values, ", "))
}

// In generates an IN expression for the specified column and the list of allowed values.
// If values is empty, a SQL "0=1" will be generated which represents a false expression.
func In(col string, values ...interface{}) Expr {
	return InExp{Col: col, Values: values, Not: false}
}

// NotIn generates an NOT IN expression for the specified column and the list of disallowed values.
// If values is empty, an empty string will be returned indicating a true expression.
func NotIn(col string, values ...interface{}) Expr {
	return InExp{Col: col, Values: values, Not: true}
}
