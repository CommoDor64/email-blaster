package pkg

import (
	"database/sql"
	"log"
)

type repo struct {
	db *sql.DB
}

type Repo interface {
	GetUsersWithLimit(startIndex int, limit int) []User
	UpdateLastSendTimestamp(int) error
}

// NewRepo return a repo struct
func NewRepo(db *sql.DB) *repo {
	return &repo{
		db: db,
	}
}

// GetUsersWithLimit selects all users from database according to a limit
// in order to provide chunks split in case the dataset is too large
func (r *repo) GetUsersWithLimit(startIndex int, limit int) []User {
	var users []User
	stmt, err := r.db.Prepare("SELECT id,firstname,lastname,email FROM user LIMIT ?,?;")
	if err != nil {
		log.Fatal(err)
	}
	rows, err := stmt.Query(startIndex, limit)
	if err != nil {
		log.Fatal(err)
	}

	for rows.Next() {
		user := User{}
		rows.Scan(&user.ID, &user.FirstName, &user.LastName, &user.Email)
		users = append(users, user)
	}

	return users
}

// UpdateLastSendTimestamp updates each user with the now-datetime
func (r *repo) UpdateLastSendTimestamp(id int) error {
	stmt, err := r.db.Prepare(`UPDATE user SET sent_at=date('now') WHERE id=?;`)
	if err != nil {
		return err
	}
	result, err := stmt.Exec(id)
	if err != nil {
		return err
	}
	if _, err = result.RowsAffected(); err != nil {
		return err
	}
	return nil
}
