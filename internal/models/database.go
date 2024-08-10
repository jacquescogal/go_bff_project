package models

import "database/sql"

type Stmt interface {
	Exec(args ...interface{}) (sql.Result, error)
	Close() error
}

type DB interface {
	Prepare(query string) (Stmt, error)
}
