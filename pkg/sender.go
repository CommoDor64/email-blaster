package pkg

import (
	"fmt"
	"log"
	"time"
)

type sender struct {
	Repo   Repo
	silent bool
}

type Sender interface {
	Send(userID int, addr string, title string, content string)
}

type Option func(s *sender)

func Silent(flag bool) func(s *sender) {
	return func(s *sender) {
		s.silent = flag
	}
}

// NewSender created a new Sender struct
func NewSender(repo Repo, done chan bool, opts ...Option) sender {

	s := sender{
		Repo: repo,
	}

	for _, opt := range opts {
		opt(&s)
	}

	return s
}

// Send is a wrapper around some email-sender client that send an email
// and updates the database accordingly.
func (s sender) Send(userID int, addr string, title string, content string) {
	err := s.sendEmail(addr, title, content)
	if err != nil {
		log.Fatal(err)
	}
	//s.Repo.UpdateLastSendTimestamp(userID)
}

// sendEmail fakes an email-sender client like SendGrid
// in this case it will always return nil as error, thus valid
func (s sender) sendEmail(addr string, title string, content string) error {
	time.Sleep(time.Millisecond * 5)
	if !s.silent {
		fmt.Printf("email sent to %s with title %s and content %s\n", addr, title, content)
	}
	return nil
}
