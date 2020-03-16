package main

import (
	"email-blaster/dev"
	"email-blaster/pkg"
	"fmt"
	"testing"
)

const (
	TestWorkerPoolSize = 1000
	TestChunkSizeTest  = 10000
	TestIsSilent       = true
	TestRowsNumber     = 100000
)

func TestIntegration(t *testing.T) {
	// setup and seed database
	dev.DatabaseSetup()
	fmt.Println("💽 seeds database")
	dev.Seed(TestRowsNumber)

	// creating all dependencies db, repository, done chan, sender client, worker pool
	fmt.Println("🤖 creating dependencies")
	db, _ := dev.NewDB()
	repo := pkg.NewRepo(db)
	done := make(chan bool)
	sender := pkg.NewSender(repo, done, pkg.Silent(TestIsSilent))
	workerPool := pkg.NewWorkerPool(done)

	// running workerpool,
	fmt.Println("🔥 starting worker goroutines")
	go workerPool.Run(TestWorkerPoolSize, func(payload interface{}) {
		p := payload.(pkg.Payload)
		sender.Send(p.UserID, p.Addr, p.Title, p.Content)
	})

	// blasting with emails
	fmt.Println("🔫 blasting emails")
	emailBlaster := pkg.NewEmailBlaster(repo, workerPool.PayloadQueue)
	emailBlaster.Blast(TestChunkSizeTest)

	// shutdown
	fmt.Println("🧯 shutdown...")
	workerPool.Shutdown()

	// remove database
	dev.DatabaseTeardown()

	// exit
	<-done
}
