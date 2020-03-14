package main

import (
	"email-blaster/dev"
	"email-blaster/pkg"
	"flag"
	"fmt"
	_ "github.com/lib/pq"
)
const (
	//TestDir = "db"
	//DBPath = "./db/eb.db"
	DefaultWorkerSizePool = 10000
	DefaultChunkSize = 100000
	DefaultIsSilent = false
	DefaultRowsNumber = 1000000
	)
var (
	workerPoolSize int
	chunkSize      int
	isSilent       bool
	rowsNumber     int
)

func main() {
	// init command line args
	flag.IntVar(&workerPoolSize, "w", DefaultWorkerSizePool, "set the amount of goroutines for worker pool")
	flag.IntVar(&chunkSize, "c", DefaultChunkSize, "set the chunk size of users from database")
	flag.BoolVar(&isSilent, "s", DefaultIsSilent, "prints to stdout the sending messages if false")
	flag.IntVar(&rowsNumber, "r", DefaultRowsNumber, "set the number of rows in seeded database")
	flag.Parse()

	// setup and seed database
	dev.DatabaseSetup()
	fmt.Println("💽 seeds database")
	dev.Seed(rowsNumber)

	// creating all dependencies db, repository, done chan, sender client, worker pool
	fmt.Println("🤖 creating dependencies")
	db, _ := dev.NewDB()
	repo := pkg.NewRepo(db)
	done := make(chan bool)
	sender := pkg.NewSender(repo, done, pkg.Silent(isSilent))
	workerPool := pkg.NewWorkerPool(done)

	// running workerpool,
	fmt.Println("🔥 starting worker goroutines")
	go workerPool.Run(workerPoolSize, func(payload interface{}) {
		p := payload.(pkg.Payload)
		sender.Send(p.UserID, p.Addr, p.Title, p.Content)
	})

	// blasting with emails
	fmt.Println("🔫 blasting emails")
	emailBlaster := pkg.NewEmailBlaster(repo, workerPool.PayloadQueue)
	emailBlaster.Blast(chunkSize)

	// shutdown
	fmt.Println("🧯 shutdown...")
	workerPool.Shutdown()

	// remove database
	dev.DatabaseTeardown()

	// exit
	<-done
}
