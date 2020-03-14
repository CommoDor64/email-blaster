package pkg

import (
	"email-blaster/dev"
	"testing"
)
const TestDir = "db"

func TestNewDB(t *testing.T) {
	db, err := dev.NewDB()
	if err != nil {
		t.Fatal(err)
	}
	db.Close()

}

func TestNewWorkerPool(t *testing.T) {
	done := make(chan bool)

	_ = NewWorkerPool(done)
}

func TestNewRepo(t *testing.T) {
	db, err := dev.NewDB()
	if err != nil {
		t.Fatal(err)
	}
	db.Close()
	_ = NewRepo(db)

}

func TestNewSender(t *testing.T) {
	db, err := dev.NewDB()
	if err != nil {
		t.Fatal(err)
	}
	db.Close()
	repo := NewRepo(db)
	done := make(chan bool)

	NewSender(repo, done)
}

func TestNewEmailBlaster(t *testing.T) {
	db, err := dev.NewDB()
	if err != nil {
		t.Fatal(err)
	}
	db.Close()
	done := make(chan bool)
	workerPool := NewWorkerPool(done)
	repo := NewRepo(db)
	_ = NewEmailBlaster(repo, workerPool.PayloadQueue)
}
