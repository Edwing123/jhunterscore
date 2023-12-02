package main

import (
	"database/sql"
	"os"

	"edwingarcia.dev/github/jhunterscore/pkg/database/sqlite"
)

func NewSQLiteDB(DSN string) *sql.DB {
	db, err := sql.Open("sqlite3", DSN)
	if err != nil {
		panic(err)
	}

	if err := db.Ping(); err != nil {
		panic(err)
	}

	return db
}

func InitDatabase(path string) error {
	err := os.WriteFile(path, sqlite.INITIALIZED_DATABASE_DATA, 0744)
	return err
}
