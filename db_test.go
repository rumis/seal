package seal

import (
	"database/sql"
	"fmt"
	"math/rand"
	"testing"
	"time"

	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/assert"
)

type User struct {
	Name       string    `seal:"name,omitempty"`
	Age        int       `seal:"age,,omitempty"`
	Class      int       `seal:"class_id"`
	CreateTime time.Time `seal:"create_time,omitempty"`
}
type UserResult struct {
	User
	ID int `seal:"id"`
}

type Class struct {
	Name string `seal:"name"`
}
type ClassResult struct {
	Class
	ID int `seal:"id"`
}

var classTableColumns = []string{"id", "name"}
var userTableColumns = []string{"id", "name", "age", "class_id", "create_time"}
var classTableColumnsNew = []string{"name"}
var userTableColumnsNew = []string{"name", "age", "class_id", "create_time"}

// get a random string
func randString(length int) string {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	bytes := make([]byte, length)
	for i := 0; i < length; i++ {
		b := r.Intn(26) + 65
		bytes[i] = byte(b)
	}
	return string(bytes)
}

// dbInit  create tables
func dbInit() (string, error) {
	dbfile := "/tmp/" + randString(10) + ".db3"

	tableClass := `CREATE TABLE class (
		id   INTEGER       PRIMARY KEY AUTOINCREMENT
						   NOT NULL
						   DEFAULT (0),
		name VARCHAR (127) NOT NULL
						   DEFAULT ""
	);`

	tableUser := `CREATE TABLE user (
		id          INTEGER       PRIMARY KEY AUTOINCREMENT
								  DEFAULT (0) 
								  NOT NULL,
		name        VARCHAR (127) NOT NULL
								  DEFAULT (''),
		age         INTEGER (10)  NOT NULL
								  DEFAULT (0),
		class_id    INTEGER (10)  NOT NULL
								  DEFAULT (0),
		create_time DATETIME      NOT NULL
								  DEFAULT (datetime('now', 'localtime') ) 
	);`

	db, err := sql.Open("sqlite3", dbfile)
	if err != nil {
		return "", err
	}
	if err := db.Ping(); err != nil {
		return "", err
	}
	defer db.Close()

	// create class table
	_, err = db.Exec(tableClass)
	if err != nil {
		return "", err
	}
	// create user table
	_, err = db.Exec(tableUser)
	if err != nil {
		return "", err
	}

	// close the db
	return dbfile, db.Close()
}

func TestDBSqlite(t *testing.T) {

	dbfile, err := dbInit()
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println("dbfile:", dbfile)

	// connect db
	db, err := Open("sqlite3", dbfile)
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()

	// INSERT
	// INSERT - struct
	c1 := Class{
		Name: "class-temp",
	}
	var newClassId int64
	err = db.Insert("class").Value(c1).Exec(&newClassId) // table:class, one row
	if err != nil {
		t.Fatal(err)
	}
	u1Name := randString(16)
	u1 := User{
		Name:       u1Name,
		Age:        13,
		Class:      int(newClassId),
		CreateTime: time.Now(),
	}
	var newUserId int64
	err = db.Insert("user").Value(u1).Exec(&newUserId) // table:user, one row
	if err != nil {
		t.Fatal(err)
	}

	// get one
	var oneCheck User
	err = db.Select("id", "name", "age", "class_id", "create_time").
		From("user").
		Where(Eq("name", u1Name)).
		Query().OneStruct(&oneCheck)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, u1Name, oneCheck.Name)

	// !u1.CreateTime.Equal(oneCheck.CreateTime) ||
	if u1.CreateTime.Format("2006-01-02 15:04:05") != oneCheck.CreateTime.Format("2006-01-02 15:04:05") {
		t.Fatal("select row time error,except:", u1.CreateTime, ", actual got:", oneCheck.CreateTime)
	}
	// assert.Equal(t, u1.CreateTime.Equal(), oneCheck)

	// multi rows
	u2Name := randString(16)
	u3Name := randString(17)
	u2 := User{
		Name:       u2Name,
		Age:        14,
		Class:      int(newClassId),
		CreateTime: time.Now(),
	}
	u3 := User{
		Name:       u3Name,
		Age:        15,
		Class:      int(newClassId),
		CreateTime: time.Now(),
	}
	var multiUserId int64
	err = db.Insert("user").Values([]User{u2, u3}).Exec(&multiUserId) // table:user, all three rows
	if err != nil {
		t.Fatal(err)
	}
	users := make([]User, 0)
	err = db.Select("name").From("user").Where(In("name", u2Name, u3Name)).OrderBy("id ASC").Query().AllStruct(&users)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, []string{u2Name, u3Name}, []string{users[0].Name, users[1].Name})
	t.Log(users)

	// INSERT - map
	// map - one
	u4Name := randString(16)
	u4 := map[string]interface{}{
		"name":        u4Name,
		"age":         16,
		"class_id":    1,
		"create_time": time.Now(),
	}
	var u4UserId int64
	err = db.Insert("user").Value(u4).Exec(&u4UserId) // table:user, all four row
	if err != nil {
		t.Fatal(err)
	}
	u4Res, err := db.Select("id", "name", "age", "class_id", "create_time").
		From("user").
		Where(Eq("id", u4UserId)).
		Query().OneMap()
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, u4["name"], u4Res["name"])

	// map - multi
	mapRows := make([]map[string]interface{}, 0)
	for i := 0; i < 10; i++ {
		tmpName := randString(11 + i)
		mapRows = append(mapRows, map[string]interface{}{
			"name":        tmpName,
			"age":         17 + i,
			"class_id":    1,
			"create_time": time.Now(),
		})
	}
	var multiMapLastRowId int64
	err = db.Insert("user").Values(mapRows).Exec(&multiMapLastRowId)
	if err != nil {
		t.Fatal(err)
	}
	var allCount int64
	err = db.Count("*").From("user").Query().Agg(&allCount)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, int64(14), allCount)

	// INSERT - column
	columns := []string{"name", "age"}
	columnValues := make([][]interface{}, 0)
	for i := 0; i < 10; i++ {
		cols := []interface{}{randString(13), i + 100}
		columnValues = append(columnValues, cols)
	}
	// column - one
	var colNewId int64
	err = db.Insert("user").Columns(columns...).Value(columnValues[0]).Exec(&colNewId)
	if err != nil {
		t.Fatal(err)
	}
	t.Log("colOneNewId:", colNewId)
	// column - multi
	err = db.Insert("user").Columns(columns...).Values(columnValues).Exec(&colNewId)
	if err != nil {
		t.Fatal(err)
	}
	t.Log("colMultiNewId:", colNewId)
	var colAllCount int64
	err = db.Count("id").From("user").Query().Agg(&colAllCount)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, int64(25), colAllCount)
	// t.Error("ttt")

	// UPDATE
	var affectCnt int64
	err = db.Update("user").Where(Eq("name", u3Name)).Value(User{
		Age: 116,
	}).Exec(&affectCnt)
	if err != nil {
		t.Fatal(err)
	}
	t.Log("update affected rows count:", affectCnt)
	var upCheckUser User
	err = db.Select("id", "name", "age").
		From("user").
		Where(Eq("name", u1Name)).
		Query().OneStruct(&upCheckUser)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, 116, upCheckUser.Age)
	// if omitempty tag is not set, the zero value will be set
	assert.Equal(t, 0, upCheckUser.Class)

	// DELETE
	var rowCountBefore int64
	err = db.Count("id").From("user").Query().Agg(&rowCountBefore)
	if err != nil {
		t.Fatal(err)
	}
	var delCount int64
	err = db.Delete("user").Where(Eq("name", u1Name)).Exec(&delCount)
	if err != nil {
		t.Fatal(err)
	}
	var rowCountAfter int64
	err = db.Count("id").From("user").Query().Agg(&rowCountAfter)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, rowCountBefore-delCount, rowCountAfter)

	// rows, err := db.Select("name", "age").
	// 	From("user").
	// 	Where(Op("name", "=", "murong")).
	// 	Query().AllMap()

	// if err != nil {
	// 	t.Fatal(err)
	// }

	// t.Error(rows)

	// u := User{}
	// err = db.Select("name", "age").
	// 	From("user").
	// 	Where(Eq("name", "murong")).
	// 	Query().OneStruct(&u)

	// if err != nil {
	// 	t.Fatal(err)
	// }
	// t.Error(u)

	// var cnt int64
	// err = db.Count("*").From("user").Query().Agg(&cnt)
	// if err != nil {
	// 	t.Fatal(err)
	// }
	// t.Error(cnt)

	// close db and remove the tmp sqlite dbfile
	closeErr := db.Close()
	if closeErr != nil {
		t.Fatal(closeErr)
	}
	// removeErr := os.Remove(dbfile)
	// if removeErr != nil {
	// 	t.Fatal(removeErr)
	// }
}

func TestOther(t *testing.T) {
	t.Error("other test")
}
