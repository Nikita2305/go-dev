package semaphore

import (
	"testing"
	"sync/atomic"
	"time"
	"math/rand"
	"sync"
)

func RandomDelay() time.Duration {
	base := 100 * time.Millisecond
	jitter := time.Duration(rand.Intn(100)-50) * time.Millisecond
	return base + jitter
}

func TestSimple(t *testing.T) {
	s := NewSemaphore(5)
	executing := atomic.Int32{}

	wg := sync.WaitGroup{}
	wg.Add(50)
	
	for i := 0; i < 50; i++ {
		go func() {
			defer wg.Done()

			s.Acquire()
			defer s.Release()
			executing.Add(1)
			if (executing.Load() > 5) {
				t.Fatal("Unexpected execution")
			}
			time.Sleep(RandomDelay())
			executing.Add(-1)
		}()
	}

	wg.Wait()
}