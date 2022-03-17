package builder

import (
	"testing"

	"github.com/rumis/seal/expr"
	"github.com/stretchr/testify/assert"
)

func TestDelete(t *testing.T) {

	b := &BuilderStandard{}
	d := NewDelete(b)

	sql, args, err := d.Table("student").Where(expr.StandardExp{
		Col:   "name",
		Op:    "=",
		Value: "murong",
	}).ToSql()
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, "DELETE FROM student WHERE name=?", sql)
	assert.Equal(t, []interface{}{"murong"}, args)

	// t.Error(sql, args)

}

func BenchmarkDeleteString(b *testing.B) {
	bs := &BuilderStandard{}
	for i := 0; i < b.N; i++ {
		d := NewDelete(bs)
		_, _, _ = d.Table("student").Where(expr.StandardExp{
			Col:   "name",
			Op:    "=",
			Value: "murong",
		}).ToSql()
	}
}
