package pkg

import "time"

type User struct {
	ID        int `json:"id" db:"id"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
	FirstName string `json:"firstname" db:"firstname"`
	LastName  string `json:"lastname" db:"lastname"`
	Email     string `json:"email" db:"email"`
	SentAt    time.Time `json:"sent_at db:"sent_at"`
	DeletedAt time.Time `json:"deleted_at", db:"deleted_at"`
}

type Payload struct {
	UserID  int    `json:"user_id"`
	Addr    string `json:"addr"`
	Title   string `json:"title"`
	Content string `json:"content"`
}
