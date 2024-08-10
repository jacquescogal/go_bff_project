package db

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

type DB struct {
	MySQL *sql.DB
}

func NewDB() (*DB, error) {
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/")
	if err != nil {
		return nil, err
	}
	return &DB{MySQL: db}, nil
}

func (db *DB) Close() error {
	return db.MySQL.Close()
}

func (db *DB) Prepare(query string) (*sql.Stmt, error) {
	return db.MySQL.Prepare(query)
}
