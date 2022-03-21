package seal

import (
	"database/sql"

	"github.com/rumis/seal/query"
)

// DB enhances sql.DB by providing a set of DB-agnostic query building methods.
// DB allows easier query building and population of data into Go variables.
type DB struct {
	query.Query

	sqlDB *sql.DB
}

// Begin starts a transaction.
func (db *DB) Begin() (*Tx, error) {
	var tx *sql.Tx
	var err error
	tx, err = db.sqlDB.Begin()
	if err != nil {
		return nil, err
	}
	return &Tx{query.NewQuery(db.Builder(), tx, db.Options()), tx}, nil
}

// Close close the db
func (db *DB) Close() error {
	return db.sqlDB.Close()
}
