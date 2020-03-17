package pkg

import (
	"fmt"
)

type emailBlaster struct {
	Repo        Repo
	PayloadQueue chan interface{}
}

type EmailBlaster interface {
	Blast(chunkSize int) error
}

// NewEmailBlaster returns an emailBlaster struct
func NewEmailBlaster(repo Repo, payloadQueue chan interface{}) *emailBlaster {
	return &emailBlaster{
		Repo:        repo,
		PayloadQueue: payloadQueue,
	}
}

// Blast reads users from the database by chunks, creates
// email payloads and messaging the worker pool via
// a dedicated channel
func (eb *emailBlaster) Blast(chunkSize int) error {
	for i := 0; ; i += chunkSize {
		if i != 0 {
			fmt.Printf("    ✅ sent %d emails\n", i)
		}
		users, err := eb.Repo.GetUsersWithLimit(i, chunkSize)
		if err != nil {
			return err
		}
		for _, user := range users {
			eb.PayloadQueue <- makePayload(user)
		}
		if len(users) < chunkSize {
			break
		}
	}
	fmt.Println("    ✅ done")
	return nil
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
