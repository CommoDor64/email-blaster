package dev

import (
	"email-blaster/pkg"
	"fmt"
	"log"
	"os"
)

const (
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
// IsFileExist
func IsFileExist(name string) bool {
	if _, err := os.Stat(name); os.IsNotExist(err) {
		return false
	}
	return true
}

// DatabaseSetup creates a directory for test database
func DatabaseSetup(testDir string) {
	err := os.Mkdir(testDir, os.ModePerm)
	if err != nil {
		log.Fatal(err)
	}
}

// DatabaseTearDown removes all
func DatabaseTeardown(testDir string) {
	err := os.RemoveAll(testDir)
	if err != nil {
		log.Fatal(err)
	}
}
func Seed(rowsNumber int) {
	db, _ := pkg.NewDB()
	db.Exec(CREATE_USER_TABLE)
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
}
