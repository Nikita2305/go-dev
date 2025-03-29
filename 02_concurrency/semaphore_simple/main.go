package main 

import (
	"fmt"
	"semaphore_simple/semaphore"
	"time"
	"math/rand"
	"sync"
)

func RandomDelay() time.Duration {
	base := 500 * time.Millisecond
	jitter := time.Duration(rand.Intn(500)-250) * time.Millisecond
	return base + jitter
}

func main() {
	s := semaphore.NewSemaphore(5)

	wg := sync.WaitGroup{}
	wg.Add(10)

	for i := 0; i < 10; i++ {
		go func(index int) {
			s.Acquire()
			defer s.Release()
			defer wg.Done()
			fmt.Println("Ready to work", index)
			time.Sleep(RandomDelay())
			fmt.Println("Work finished!!", index)
		}(i)
	}

	wg.Wait()
}