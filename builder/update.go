package builder

import (
	"errors"

	"github.com/rumis/seal/expr"
	"github.com/rumis/seal/utils"
)

// Update represents a UPDATE query
// It can be built into a update sql clauses and params by calling the ToSql method.
type Update struct {
	b Builder

	table string
	val   map[string]interface{}
	where expr.Expr
}

// NewUpdate 创建新更新器
func NewUpdate(b Builder) *Update {
	return &Update{
		b: b,
	}
}

// Into specifies which tables to insert into.
func (u *Update) Table(table string) *Update {
	u.table = table
	return u
}

// Value specifies the update column and value
// type of val can be map[string]interface{}, struct
func (u *Update) Value(val interface{}) *Update {
	if v, ok := val.(map[string]interface{}); ok {
		u.val = v
		return u
	}
	v, err := utils.Struct2Map(val)
	if err != nil {
		return u
	}
	u.val = v
	return u
}

// Where specifies the WHERE condition.
func (u *Update) Where(e expr.Expr) *Update {
	if u.where == nil {
		u.where = e
		return u
	}
	return u.AndWhere(e)
}

// AndWhere concatenates a new WHERE condition with the existing one (if any) using "AND".
func (u *Update) AndWhere(e expr.Expr) *Update {
	u.where = expr.AndOrExp{
		Exps: []expr.Expr{u.where, e},
		Op:   "AND",
	}
	return u
}

// OrWhere concatenates a new WHERE condition with the existing one (if any) using "OR".
func (u *Update) OrWhere(e expr.Expr) *Update {
	u.where = expr.AndOrExp{
		Exps: []expr.Expr{u.where, e},
		Op:   "OR",
	}
	return u
}

// ToSql build the sql clauses and params
func (u *Update) ToSql() (string, []interface{}, error) {
	if u.val == nil {
		return "", nil, errors.New("update value not set")
	}
	if u.table == "" {
		return "", nil, errors.New("table name not set")
	}
	if u.where == nil {
		return "", nil, errors.New("update should have a where clauses")
	}
	params := expr.Params{}

	sql := u.b.Update(u.table, u.val, params) + " " + u.b.Where(u.where, params)

	return utils.ReplacePlaceHolders(sql, u.b.Placeholder(), params)
}
