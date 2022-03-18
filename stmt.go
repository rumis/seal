package seal

import (
	"github.com/rumis/seal/expr"
)

// Eq generates a Standard equal expression
func Eq(col string, val interface{}) expr.Expr {
	return expr.Op(col, "=", val)
}

// StaticEq generates a static equal expression which without params
func StaticEq(col1 string, col2 string) expr.Expr {
	return StaticOp(col1, "=", col2)
}

// Op generates a Standard expression
// eg. >,>=,<,<=,....
func Op(col string, op string, val interface{}) expr.Expr {
	return expr.Op(col, op, val)
}

// StaticOp generates a Standard expression
// Which always used for multi table select
func StaticOp(col1 string, op string, col2 string) expr.Expr {
	return expr.New(col1 + op + col2)
}

// Not generates a NOT expression which prefixes "NOT" to the specified expression.
func Not(e expr.Expr) expr.Expr {
	return expr.Not(e)
}

// And generates an AND expression which concatenates the given expressions with "AND".
func And(exps ...expr.Expr) expr.Expr {
	if len(exps) == 1 {
		return exps[0]
	}
	return expr.Group("AND", exps...)
}

// Or generates an OR expression which concatenates the given expressions with "OR".
func Or(exps ...expr.Expr) expr.Expr {
	if len(exps) == 1 {
		return exps[0]
	}
	return expr.Group("OR", exps...)
}

// In generates an IN expression for the specified column and the list of allowed values.
// If values is empty, a SQL "0=1" will be generated which represents a false expression.
func In(col string, values ...interface{}) expr.Expr {
	return expr.In(col, values...)
}

// NotIn generates an NOT IN expression for the specified column and the list of disallowed values.
// If values is empty, an empty string will be returned indicating a true expression.
func NotIn(col string, values ...interface{}) expr.Expr {
	return expr.NotIn(col, values...)
}

// Like generates a LIKE expression for the specified column and the possible strings that the column should be like.
// If multiple values are present, the column should be like *all* of them. For example, Like("name", "key", "word")
// will generate a SQL expression: "name" LIKE "%key%" AND "name" LIKE "%word%".
func Like(col string, value string) expr.LikeExp {
	return expr.Like(col, value)
}

// NotLike generates a NOT LIKE expression.
// For example, NotLike("name", "key", "word") will generate a SQL expression:
// "name" NOT LIKE "%key%" AND "name" NOT LIKE "%word%". Please see Like() for more details.
func NotLike(col string, value string) expr.LikeExp {
	return expr.NotLike(col, value)
}

// Exists generates an EXISTS expression by prefixing "EXISTS" to the given expression.
func Exists(exp expr.Expr) expr.Expr {
	return expr.Exists(exp)
}

// NotExists generates an EXISTS expression by prefixing "NOT EXISTS" to the given expression.
func NotExists(exp expr.Expr) expr.Expr {
	return expr.NotExists(exp)
}

// Between generates a BETWEEN expression.
// For example, Between("age", 10, 30) generates: "age" BETWEEN 10 AND 30
func Between(col string, from, to interface{}) expr.Expr {
	return expr.Between(col, from, to)
}

// NotBetween generates a NOT BETWEEN expression.
// For example, NotBetween("age", 10, 30) generates: "age" NOT BETWEEN 10 AND 30
func NotBetween(col string, from, to interface{}) expr.Expr {
	return expr.NotBetween(col, from, to)
}

// Count generates a COUNT() expression
// For example:
//	Count("id"): 					Count(id)
//	Count("id","id_count"): 		Count(id) AS id_count
//	Count("id","id_count","user"): 	Count(user.id) AS id_count
func Count(col string, alias_table ...string) expr.Expr {
	return expr.Aggregate("COUNT", col, alias_table...)
}

// SUM generates a SUM() expression
func Sum(col string, alias_table ...string) expr.Expr {
	return expr.Aggregate("SUM", col, alias_table...)
}

// Count generates a MAX() expression
func Max(col string, alias_table ...string) expr.Expr {
	return expr.Aggregate("MAX", col, alias_table...)
}

// Min generates a MIN() expression
func Min(col string, alias_table ...string) expr.Expr {
	return expr.Aggregate("MIN", col, alias_table...)
}

// Avg generates an AVG() expression
func Avg(col string, alias_table ...string) expr.Expr {
	return expr.Aggregate("AVG", col, alias_table...)
}
