package query

import "database/sql"

// Result enhances sql.Result by providing additional exec methods
type Result struct {
	sr  sql.Result
	err error
}

// NewExecResult generates an sql.Result instance with *sql.Result and error which come from db.Exec
func NewExecResult(r sql.Result, err error) Result {
	return Result{
		sr:  r,
		err: err,
	}
}

//  LastInsertId return the last new rows id
func (r Result) LastInsertId() (int64, error) {
	if r.err != nil {
		return 0, r.err
	}
	return r.sr.LastInsertId()
}

// RowsAffected return how much rows affected
func (r Result) RowsAffected() (int64, error) {
	if r.err != nil {
		return 0, r.err
	}
	return r.sr.RowsAffected()
}
