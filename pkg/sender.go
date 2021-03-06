package pkg

import (
	"fmt"
	"time"
)

type sender struct {
	Repo           Repo
	silent         bool
	updateDatabase bool
	latency        time.Duration
}

type Sender interface {
	Send(userID int, addr string, title string, content string) error
}

type Option func(s *sender)

// Slient sets whether the function should print
// debug output
func Silent(flag bool) func(s *sender) {
	return func(s *sender) {
		s.silent = flag
	}
}

// UpdateDatabase sets whether the database should updated
// with the corresponding send datetime
func UpdateDatabase(flag bool) func(s *sender) {
	return func(s *sender) {
		s.updateDatabase = flag
	}
}

// Latency sets the latency config field
func Latency(latency time.Duration) func(s *sender) {
	return func(s *sender) {
		s.latency = latency
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
func (s sender) Send(userID int, addr string, title string, content string) error {
	err := s.sendEmail(addr, title, content)
	if err != nil {
		return err
	}
	if !s.updateDatabase {
		return nil
	}

	err = s.Repo.UpdateLastSendTimestamp(userID)
	if err != nil {
		return err
	}
	return nil
}

// sendEmail fakes an email-sender client like SendGrid
// in this case it will always return nil as error, thus valid
func (s sender) sendEmail(addr string, title string, content string) error {
	time.Sleep(time.Millisecond * s.latency)
	if !s.silent {
		fmt.Printf("email sent to %s with title %s and content %s\n", addr, title, content)
	}
	return nil
}
