package builder

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInsertStruct(t *testing.T) {

	b := &BuilderStandard{}
	in := NewInsert(b)

	sql, arg, err := in.Into("student").Value(struct {
		Name string `seal:"name"`
		Age  int    `seal:"age"`
	}{
		Name: "murong",
		Age:  13,
	}).ToSql()

	if err != nil {
		t.Fatal(err)
	}

	assert.Contains(t, []string{
		"INSERT INTO student (name, age) VALUES (?,?)",
		"INSERT INTO student (age, name) VALUES (?,?)",
	}, sql)
	assert.Contains(t, [][]interface{}{{"murong", 13}, {13, "murong"}}, arg)

	// t.Error(sql, arg)

	in1 := NewInsert(b)
	sql1, arg1, err1 := in1.Into("student").Values([]struct {
		Name string `seal:"name,omitvalue:murong"`
		Age  int    `seal:"age"`
	}{{
		Name: "murong",
		Age:  13,
	}, {
		Name: "murong",
		Age:  14,
	}}).ToSql()

	if err1 != nil {
		t.Fatal(err1)
	}

	assert.Equal(t, "INSERT INTO student (age) VALUES (?), (?)", sql1)
	assert.Contains(t, [][]interface{}{{13, 14}, {14, 13}}, arg1)

	// t.Error(sql1, arg1)

}

func TestInsertMap(t *testing.T) {

	b := &BuilderStandard{}
	in := NewInsert(b)

	sql, arg, err := in.Into("student").Value(map[string]interface{}{
		"name": "murong",
		"age":  13,
	}).ToSql()

	if err != nil {
		t.Fatal(err)
	}

	assert.Contains(t, []string{
		"INSERT INTO student (name, age) VALUES (?,?)",
		"INSERT INTO student (age, name) VALUES (?,?)",
	}, sql)
	assert.Contains(t, [][]interface{}{{"murong", 13}, {13, "murong"}}, arg)

	// t.Error(sql, arg)

	in1 := NewInsert(b)
	sql1, arg1, err1 := in1.Into("student").Values([]map[string]interface{}{{
		"age": 13,
	}, {
		"age": 14,
	}}).ToSql()

	if err1 != nil {
		t.Fatal(err1)
	}

	assert.Equal(t, "INSERT INTO student (age) VALUES (?), (?)", sql1)
	assert.Contains(t, [][]interface{}{{13, 14}, {14, 13}}, arg1)

	// t.Error(sql1, arg1)

}

func TestInsertValue(t *testing.T) {

	b := &BuilderStandard{}
	in := NewInsert(b)

	sql, arg, err := in.Into("student").Columns("name", "age").Value([]interface{}{"murong", 13}).ToSql()

	if err != nil {
		t.Fatal(err)
	}

	assert.Contains(t, []string{
		"INSERT INTO student (name, age) VALUES (?,?)",
	}, sql)
	assert.Contains(t, [][]interface{}{{"murong", 13}}, arg)

	// t.Error(sql, arg)

	in1 := NewInsert(b)
	sql1, arg1, err1 := in1.Into("student").Columns("name", "age").Values([][]interface{}{
		{"murong", 13},
		{"liu", 14},
	}).ToSql()

	if err1 != nil {
		t.Fatal(err1)
	}

	assert.Equal(t, "INSERT INTO student (name, age) VALUES (?,?), (?,?)", sql1)
	assert.Contains(t, [][]interface{}{{"murong", 13, "liu", 14}}, arg1)

	// t.Error(sql1, arg1)

}
