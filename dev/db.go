package dev

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"os"
	"strings"
)

const (
	DBString          = "./db/eb.db?cache=shared&mode=rwc&_journal_mode=WAL"
	DBPath            = "./db/eb.db"
	TestDir           = "./db"
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

type insertPayload struct {
	Firstname string
	Lastname  string
	Email     string
}

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

var insert = "INSERT INTO user (firstname,lastname,email) VALUES %s"

const insertBatchSize = 100

func Seed(rowsNumber int) {
	db, _ := NewDB()
	// create table
	db.Exec(CREATE_USER_TABLE)

	var stmt string
	var stmtValues []string
	var stmtArgs []interface{}
	indexInBatch := 0
	for i := 0; i < rowsNumber; i++ {
		if i%(insertBatchSize) == 0 && i != 0 {
			stmt = fmt.Sprintf(insert, strings.Join(stmtValues, ","))
			_, err := db.Exec(stmt, stmtArgs...)
			if err != nil {
				log.Fatal(err)
			}
			stmtValues = stmtValues[:0]
			indexInBatch = 0
			stmtArgs = stmtArgs[:0]
		}
		if i%100000 == 0 && i != 0 {
			fmt.Printf("    ✅ %d records written\n", i)
		}
		stmtValues = append(stmtValues, fmt.Sprintf("($%d,$%d,$%d)",
			indexInBatch*3+1,
			indexInBatch*3+2,
			indexInBatch*3+3))

		stmtArgs = append(stmtArgs,
			fmt.Sprintf("dor%d", i),
			fmt.Sprintf("cohen%d", i),
			fmt.Sprintf("dorcohen%d@gmail.com", i))

		indexInBatch++

	}
	fmt.Println("    ✅ done")
}