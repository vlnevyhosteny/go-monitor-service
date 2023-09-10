package database

import (
	"database/sql"
	"sync"

	_ "github.com/mattn/go-sqlite3"
)

var (
	connection     *sql.DB
	connectionOnce sync.Once
)

func NewConnection(databaseFile string) (*sql.DB, error) {
	return sql.Open("sqlite3", databaseFile)
}

func ProvideConnection(databaseFile string) *sql.DB {
	connectionOnce.Do(func() {
		var err error
		connection, err = NewConnection(databaseFile)
		if err != nil {
			panic(err)
		}
	})

	return connection
}
