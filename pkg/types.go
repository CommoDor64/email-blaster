package pkg

type User struct {
	ID        int `json:"id" db:"id"`
	UUID      string `json:"uuid", db:"uuid"`
	CreatedAt string `json:"created_at" db:"created_at"`
	UpdatedAt string `json:"updated_at" db:"updated_at"`
	FirstName string `json:"firstname" db:"firstname"`
	LastName  string `json:"lastname" db:"lastname"`
	Email     string `json:"email" db:"email"`
	SentAt    string `json:"sent_at db:"sent_at"`
	DeletedAt string `json:"deleted_at", db:"deleted_at"`
}

type Payload struct {
	UserID  int    `json:"user_id"`
	Addr    string `json:"addr"`
	Title   string `json:"title"`
	Content string `json:"content"`
}
