package pkg

import (
	"fmt"
)

type emailBlaster struct {
	Repo        Repo
	WorkerQueue chan interface{}
}

type EmailBlaster interface {
	Blast(chunkSize int)
}

// NewEmailBlaster returns an emailBlaster struct
func NewEmailBlaster(repo Repo, workerQueue chan interface{}) *emailBlaster {
	return &emailBlaster{
		Repo:        repo,
		WorkerQueue: workerQueue,
	}
}

// Blast reads users from the database by chunks, creates
// email payloads and messaging the worker pool via
// a dedicated channel
func (eb *emailBlaster) Blast(chunkSize int) {
	for i := 0; ; i += chunkSize {
		if i != 0 {
			fmt.Printf("    ✅ sent %d emails\n", i)
		}
		users := eb.Repo.GetUsersWithLimit(i, chunkSize)
		for _, user := range users {
			eb.WorkerQueue <- makePayload(user)
		}
		if len(users) < chunkSize {
			break
		}
	}
}

// makePayload is a convenience function that translates
// the user database model/type to a business logic payload
// model/type
func makePayload(user User) Payload {
	return Payload{
		UserID:  user.ID,
		Addr:    user.Email,
		Title:   fmt.Sprintf("hello %s %s", user.FirstName, user.LastName),
		Content: "lorem ipsum...",
	}
}
