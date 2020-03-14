package pkg

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
)

// NewDB returns a new go sql wrapper
func NewDB() (*sql.DB, error) {
	db, err := sql.Open("sqlite3", "./test/eb.db?cache=shared&mode=rwc&_journal_mode=WAL")
	if err != nil {
		return nil, err
	}
	return db, nil
}
