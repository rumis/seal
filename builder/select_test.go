package builder

import (
	"testing"

	"github.com/rumis/seal/expr"
	"github.com/stretchr/testify/assert"
)

func TestSelect1(t *testing.T) {

	b := &BuilderStandard{}
	s := NewSelect(b)

	sql, args, err := s.Select("name", "age").
		From("student").
		Where(expr.In("age", 13, 14)).
		AndWhere(expr.Like("name", `%mu%`)).
		OrWhere(expr.Group("AND", expr.Op("age", ">", 100), expr.Op("age", "<", 200))).
		ToSql()

	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, "SELECT name,age FROM student WHERE age IN (?, ?) AND name LIKE ? OR (age>? AND age<?)",
		sql)

	assert.Equal(t, []interface{}{13, 14, "%mu%", 100, 200}, args)

	// t.Error(sql, args)

}

func TestJoin1(t *testing.T) {

	b := &BuilderStandard{}
	s := NewSelect(b)

	_, _, err := s.Select("name as nx", "age").
		From("student as c").
		Where(expr.In("age", 13, 14)).InnerJoin("school as s", expr.Exp{
		E: "student.name=school.name",
	}, "t1", "c2").
		ToSql()

	if err != nil {
		t.Fatal(err)
	}
	// t.Error(sql, args)

	s1 := NewSelect(b)
	sql1, args1, err := s1.Select("id as coursewareid").
		AndSelect("cp", "id as packageid").
		From("xes_coursewares as c", "xes_courseware_relation_package as cp", "xes_page_packages as pack").
		Where(expr.New("c.id=cp.courseware_id")).
		Where(expr.Op("c.id", "=", 10086)).
		ToSql()
	if err != nil {
		t.Fatal(err)
	}
	esql1 := "SELECT c.id as coursewareid,cp.id as packageid FROM xes_coursewares as c, xes_courseware_relation_package as cp, xes_page_packages as pack WHERE c.id=cp.courseware_id AND c.id=?"

	// t.Error(sql1, args1)

	assert.Equal(t, esql1, sql1)
	assert.Equal(t, []interface{}{10086}, args1)

}

func TestAggregate(t *testing.T) {

	b := &BuilderStandard{}
	s := NewSelect(b)

	sql, args, err := s.Agg("COUNT", "id", "agg_count").
		From("student").
		Where(expr.In("age", 13, 14)).
		ToSql()

	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, "SELECT COUNT(id) AS agg_count FROM student WHERE age IN (?, ?)", sql)

	assert.Equal(t, []interface{}{13, 14}, args)

}
