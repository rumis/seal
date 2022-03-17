package seal

import (
	"testing"

	_ "github.com/mattn/go-sqlite3"
)

func TestDB(t *testing.T) {

	db, _ := Open("mysql", "")

	err := db.Update("xx").Value(1)

	if err != nil {
		t.Fatal(err)
	}
}

type User struct {
	Name string `seal:"name"`
	Age  int    `seal:"age"`
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

}
