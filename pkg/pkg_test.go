package pkg

import (
	"testing"
)

func TestSend(t *testing.T) {
	Send("commo64dor@gmail.com", "hello world")
}

//func TestRepo_GetAllUsers(t *testing.T) {
//	db, _ := NewDB()
//	repo := NewRepo(db)
//	repo.GetAllUsers()
//}

//func TestSpawnWorkerNodes(t *testing.T) {
//	wg := sync.WaitGroup{}
//	ch := make(chan Payload)
//	wg.Add(10)
//	SpawnWorkerNodes(1000, ch, wg)
//
//	for i := 0; i < 10; i++ {
//		payload := Payload{
//			addr:    fmt.Sprintf("dor%d", i),
//			content: fmt.Sprintf("hello%d", i),
//		}
//		ch <- payload
//	}
//
//}

//func TestSeed(t *testing.T) {
//	Seed()
//}
