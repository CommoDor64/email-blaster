package dev

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"fmt"
	"log"
	"os"
)

const (
	DBString = "./db/eb.db?cache=shared&mode=rwc&_journal_mode=WAL"
	DBPath = "./db/eb.db"
	TestDir="./db"
	CREATE_USER_TABLE = `CREATE TABLE IF NOT EXISTS user (
					id  INTEGER PRIMARY KEY AUTOINCREMENT,
					created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
					updated_at DATETIME DEFAULT CURRENT_TIMESTAMP,
					firstname TEXT,
					lastname TEXT,
					email TEXT,
					sent_at DATETIME,
					deleted_at
					);`
)

func NewDB() (*sql.DB, error) {
	db, err := sql.Open("sqlite3", DBString)
	if err != nil {
		return nil, err
	}
	return db, nil
}

// IsFileExist
func IsFileExist(name string) bool {
	if _, err := os.Stat(name); os.IsNotExist(err) {
		return false
	}
	return true
}

// DatabaseSetup creates a directory for test database
func DatabaseSetup() {
	_ = os.Mkdir(TestDir, os.ModePerm)
}

// DatabaseTearDown removes all
func DatabaseTeardown() {
	err := os.RemoveAll(TestDir)
	if err != nil {
		log.Fatal(err)
	}
}
func Seed(rowsNumber int) {
	db, _ := NewDB()
	// create table
	db.Exec(CREATE_USER_TABLE)

	// seed user table
	stmt, err := db.Prepare("INSERT INTO user (firstname,lastname,email) VALUES (?,?,?)")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	for i := 0; i < rowsNumber; i++ {
		if i%(100000) == 0 && i != 0 {
			fmt.Printf("    ✅ %d records written\n", i)
		}
		stmt.Exec(
			fmt.Sprintf("%s%d", "dor", i),
			fmt.Sprintf("%s%d", "cohen", i),
			fmt.Sprintf("%s%d@gmail.com", "dor", i))
	}
	fmt.Println("    ✅ done")
}
