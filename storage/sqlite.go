package storage

import (
	"database/sql"
	_"github.com/mattn/go-sqlite3"
)

type SqliteDriver struct {
	DB *sql.DB
}

func init() {
	RegisterDriver("sqlite3", &SqliteDriver{})
}

func (sd *SqliteDriver) Open(connectionString string) error {
	db, err := sql.Open("sqlite3", connectionString)
	if err != nil {
		return err
	} else {
		sd.DB = db
		return nil
	}
}
