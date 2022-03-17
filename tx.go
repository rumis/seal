package seal

import (
	"database/sql"

	"github.com/rumis/seal/query"
)

// Tx enhances sql.Tx with additional querying methods.
type Tx struct {
	query.Query

	tx *sql.Tx
}

// Commit commits the transaction.
func (t *Tx) Commit() error {
	return t.tx.Commit()
}

// Rollback aborts the transaction.
func (t *Tx) Rollback() error {
	return t.tx.Rollback()
}
