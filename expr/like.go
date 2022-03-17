package expr

import (
	"fmt"
)

// LikeExp represents a variant of LIKE expressions.
type LikeExp struct {
	Col   string
	Value string
	// Like stores the LIKE operator. It can be "LIKE", "NOT LIKE".
	// It may also be customized as something like "ILIKE".
	Like string
}

// Build converts an expression into a SQL fragment.
func (e LikeExp) Build(params Params) string {
	if e.Value == "" {
		return ""
	}
	key := fmt.Sprintf("p%v", len(params))
	params[key] = e.Value
	return fmt.Sprintf("%v %v {:%v}", e.Col, e.Like, key)
}

// Like generates a LIKE expression for the specified column and the possible strings that the column should be like.
// If multiple values are present, the column should be like *all* of them. For example, Like("name", "key", "word")
// will generate a SQL expression: "name" LIKE "%key%" AND "name" LIKE "%word%".
func Like(col string, value string) LikeExp {
	return LikeExp{
		Col:   col,
		Value: value,
		Like:  "LIKE",
	}
}

// NotLike generates a NOT LIKE expression.
// For example, NotLike("name", "key", "word") will generate a SQL expression:
// "name" NOT LIKE "%key%" AND "name" NOT LIKE "%word%". Please see Like() for more details.
func NotLike(col string, value string) LikeExp {
	return LikeExp{
		Col:   col,
		Value: value,
		Like:  "NOT LIKE",
	}
}
