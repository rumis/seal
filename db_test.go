package seal

import (
	"testing"

	_ "github.com/mattn/go-sqlite3"
)

type User struct {
	ID         int    `seal:"id"`
	Name       string `seal:"name"`
	Age        int    `seal:"age"`
	Class      int    `seal:"class_id"`
	CreateTime int    `seal:"create_time"`
}

type Class struct {
	ID   int    `seal:"id"`
	Name string `seal:"name"`
}

func TestDBSqlite(t *testing.T) {

	dbfile := "/home/ubuntu/workspace/data/seal.db3"

	db, err := Open("sqlite3", dbfile)
	if err != nil {
		t.Fatal(err)
	}

	rows, err := db.Select("name", "age").
		From("user").
		Where(Op("name", "=", "murong")).
		Query().AllMap()

	if err != nil {
		t.Fatal(err)
	}

	t.Error(rows)

	u := User{}
	err = db.Select("name", "age").
		From("user").
		Where(Eq("name", "murong")).
		Query().OneStruct(&u)

	if err != nil {
		t.Fatal(err)
	}
	t.Error(u)

	var cnt int64
	err = db.Count("*").From("user").Query().Agg(&cnt)
	if err != nil {
		t.Fatal(err)
	}
	t.Error(cnt)
}
