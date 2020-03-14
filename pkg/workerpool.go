package pkg

import "sync"

type workerPool struct {
	JobQueue chan interface{}
	Done chan bool
	wg sync.WaitGroup
}
type WorkerPool interface {
	Run(workerPoolSize int, job Job)
	Spawn(job Job)
	Shutdown()
}

// Job is a function type that describes operation of each
// worker, since it is a workerpool, all workers are doing
// the same thing
type Job func(interface{})

// NewWorkerPool returns a struct representing a workerpool
func NewWorkerPool(done chan bool) *workerPool {
	return &workerPool{
		JobQueue: make(chan interface{}),
		Done: done,
		wg: sync.WaitGroup{},
	}
}

// Run fires up all workers according to received number
// and Job description
func (wp *workerPool) Run(workerPoolSize int, job Job) {
	wp.wg.Add(workerPoolSize)
	for i := 0; i < workerPoolSize; i++ {
		go wp.Spawn(job)
	}
	wp.wg.Wait()
	wp.Done <- true
}

// Spawn fires up a single worker
func (wp *workerPool) Spawn(job Job) {
	for payload := range wp.JobQueue {
		job(payload)
	}
	wp.wg.Done()
}

// Shutdown terminates the workerpool
func (wp *workerPool) Shutdown() {
	close(wp.JobQueue)
}
