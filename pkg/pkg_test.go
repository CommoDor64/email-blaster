package pkg

import (
	"email-blaster/dev"
	"testing"
)

// DB
func TestNewDB(t *testing.T) {
	db, err := dev.NewDB()
	if err != nil {
		t.Fatal(err)
	}
	db.Close()

}

// WorkerPool
func TestNewWorkerPool(t *testing.T) {
	done := make(chan bool)

	_ = NewWorkerPool(done)
}

func TestWorkerPool_RunShutDown(t *testing.T) {
	done := make(chan bool)

	workerPool := NewWorkerPool(done)
	job := func(interface{}) {}
	go workerPool.Run(10, job)
	workerPool.Shutdown()
}

// Repo
func TestNewRepo(t *testing.T) {
	db, err := dev.NewDB()
	if err != nil {
		t.Fatal(err)
	}
	db.Close()
	_ = NewRepo(db)

}

// Sender
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

// EmailBlaster
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
