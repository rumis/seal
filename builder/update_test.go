package builder

import (
	"testing"

	"github.com/rumis/seal/expr"
	"github.com/stretchr/testify/assert"
)

type Student struct {
	Name string `seal:"name"`
	Age  int    `seal:"age,omitvalue:-1"`
}

func TestUpdateStruct(t *testing.T) {
	b := &BuilderStandard{}
	u := NewUpdate(b)

	// simple where
	sql, params, err := u.Table("student").Value(Student{
		Name: "murong",
		Age:  13,
	}).Where(&expr.StandardExp{
		Col:   "age",
		Op:    "=",
		Value: 13,
	}).ToSql()

	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, "UPDATE student SET name=?, age=? WHERE age=?", sql)
	assert.Equal(t, []interface{}{"murong", 13, 13}, params)

	// where and
	sql1, params1, err1 := u.Table("student").Value(Student{
		Name: "murong",
		Age:  13,
	}).Where(expr.StandardExp{
		Col:   "age",
		Op:    "=",
		Value: 13,
	}).AndWhere(expr.StandardExp{
		Col:   "name",
		Op:    "=",
		Value: "murong",
	}).ToSql()

	if err1 != nil {
		t.Fatal(err)
	}
	assert.Contains(t, []string{
		"UPDATE student SET age=?, name=? WHERE age=? AND name=?",
		"UPDATE student SET name=?, age=? WHERE age=? AND name=?"}, sql1)
	assert.Contains(t, [][]interface{}{{"murong", 13, 13, "murong"}, {13, "murong", 13, "murong"}}, params1)

	// where and or ,change priority

	sql2, params2, err2 := u.Table("student").Value(Student{
		Name: "murong",
		Age:  -1,
	}).Where(expr.StandardExp{
		Col:   "age",
		Op:    "=",
		Value: 13,
	}).AndWhere(expr.GroupExp{
		Exps: []expr.Expr{
			expr.StandardExp{
				Col:   "age",
				Op:    "=",
				Value: 13,
			},
			expr.StandardExp{
				Col:   "name",
				Op:    "=",
				Value: "murong",
			},
		},
		Op: "OR",
	}).ToSql()
	if err2 != nil {
		t.Fatal(err2)
	}
	assert.Equal(t, "UPDATE student SET name=? WHERE age=? AND (age=? OR name=?)", sql2)
	assert.Equal(t, []interface{}{"murong", 13, 13, "murong"}, params2)

	// t.Error(sql2, params2, err2)
}

func BenchmarkUpdateStruct(b *testing.B) {

	builder := &BuilderStandard{}
	u := NewUpdate(builder)
	for i := 0; i < b.N; i++ {
		_, _, _ = u.Table("student").Value(Student{
			Name: "murong",
			Age:  13,
		}).Where(&expr.StandardExp{
			Col:   "age",
			Op:    "=",
			Value: 13,
		}).ToSql()
	}
}

func TestUpdateMap(t *testing.T) {

	b := &BuilderStandard{}
	u := NewUpdate(b)

	// simple where
	sql, params, err := u.Table("student").Value(map[string]interface{}{
		"name": "murong",
		"age":  13,
	}).Where(&expr.StandardExp{
		Col:   "age",
		Op:    "=",
		Value: 13,
	}).ToSql()

	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, "UPDATE student SET name=?, age=? WHERE age=?", sql)
	assert.Equal(t, []interface{}{"murong", 13, 13}, params)

}
