package mysqlt

import (
	"database/sql"
)

// SQLTemplate define database operation
type SQLTemplate interface {
	Insert(query string, uid *int, args ...interface{}) error
	Select(query string, callback RowMapperCallback) error
	SelectOne(query string, callback RowMapperCallback, args ...interface{}) error
}

// RowMapperCallback callback function
type RowMapperCallback func(RowScanner) error

// RowScanner define scan function
type RowScanner interface {
	Scan(dest ...interface{}) error
}

// interface for supporting sql.DB and sql.Tx to do sql statement
type sqlExecer interface {
	Exec(query string, args ...interface{}) (sql.Result, error)
	Query(query string, args ...interface{}) (*sql.Rows, error)
	Prepare(query string) (*sql.Stmt, error)
}

// do database operation
type operate struct {
	sqlExecer
}

// Insert implement Insert function defined by SQLTemplate interface
func (o *operate) Insert(query string, uid *int, args ...interface{}) error {
	stmt, err := o.Prepare(query)
	checkErr(err)
	defer stmt.Close()

	res, err := o.Exec(query, args...)
	checkErr(err)
	if err != nil {
		return err
	}
	id, err := res.LastInsertId()
	if err != nil {
		return err
	}
	*uid = int(id)

	return err
}

func (o *operate) Select(query string, callback RowMapperCallback) error {
	rows, err := o.Query(query)
	checkErr(err)
	defer rows.Close()

	for rows.Next() {
		err := callback(rows)
		checkErr(err)
		if err != nil {
			return err
		}
	}

	return err
}

func (o *operate) SelectOne(query string, callback RowMapperCallback, args ...interface{}) error {
	stmt, err := o.Prepare(query)
	checkErr(err)
	defer stmt.Close()

	row := stmt.QueryRow(args...)
	err = callback(row)
	checkErr(err)

	return err
}

// NewSQLTemplate create a new template and return
func NewSQLTemplate(db sqlExecer) SQLTemplate {
	o := &operate{db}
	return o
}

// check error
func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
