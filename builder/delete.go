package builder

import (
	"github.com/rumis/seal/expr"
	"github.com/rumis/seal/utils"
)

// import "github.com/rumis/seal/expr"

// Delete represents a DELETE query
// It can be built into a delete sql clauses and params by calling the ToSql method.
type Delete struct {
	b Builder

	table string
	where expr.Expr
}

// NewDelete
func NewDelete(b Builder) *Delete {
	return &Delete{
		b: b,
	}
}

// Table specifies which tables to insert into.
func (d *Delete) Table(table string) *Delete {
	d.table = table
	return d
}

// Where specifies the WHERE condition.
func (d *Delete) Where(e expr.Expr) *Delete {
	if d.where == nil {
		d.where = e
	}
	return d.AndWhere(e)
}

// AndWhere concatenates a new WHERE condition with the existing one (if any) using "AND".
func (d *Delete) AndWhere(e expr.Expr) *Delete {
	d.where = expr.AndOrExp{
		Exps: []expr.Expr{d.where, e},
		Op:   "AND",
	}
	return d
}

// OrWhere concatenates a new WHERE condition with the existing one (if any) using "OR".
func (d *Delete) OrWhere(e expr.Expr) *Delete {
	d.where = expr.AndOrExp{
		Exps: []expr.Expr{d.where, e},
		Op:   "OR",
	}
	return d
}

// ToSql build the sql clauses and params
func (d *Delete) ToSql() (string, []interface{}, error) {
	params := expr.Params{}
	sql := d.b.Delete(d.table) + " " + d.b.Where(d.where, params)
	return utils.ReplacePlaceHolders(sql, d.b.Placeholder(), params)
}
