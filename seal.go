package seal

import (
	"database/sql"

	"github.com/rumis/seal/builder"
	"github.com/rumis/seal/query"
)

// Open opens a database specified by a driver name and data source name (DSN).
// Note that Open does not check if DSN is specified correctly. It doesn't try to establish a DB connection either.
// Please refer to sql.Open() for more information.
func Open(driverName string, sourceName string) (DB, error) {
	db, err := sql.Open(driverName, sourceName)
	if err != nil {
		return DB{}, err
	}
	if err := db.Ping(); err != nil {
		return DB{}, err
	}
	var b builder.Builder
	switch driverName {
	case "mysql":
		b = builder.NewMysqlBuilder()
	case "sqlite3":
		b = builder.NewSqliteBuilder()
	default:
		b = builder.NewStandardBuilder()
	}
	return DB{
		query.NewQuery(b, db),
		db,
	}, nil
}
