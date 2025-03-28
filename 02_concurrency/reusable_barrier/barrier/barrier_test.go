package barrier

import (
	"testing"
	"time"
	"sync"
	"sync/atomic"
)

func TestSimple(t *testing.T) {
	n := 5
	b := NewBarrier(int32(n))
	executed := atomic.Int32{}
	for i := 0; i < n - 1; i++ {
		go func() {
			b.Wait()
			executed.Add(1)
		}()
	}

	time.Sleep(time.Millisecond * 200)

	if (executed.Load() != 0) {
		t.Fatal("Executed too early")
	}

	b.Wait()
	executed.Add(1)

	time.Sleep(time.Millisecond * 200)
	if (executed.Load() != int32(n)) {
		t.Fatal("Didn't execute at all")
	}
}

func TestCyclic(t *testing.T) {
	epochs := 10
	n := 5
	b := NewBarrier(int32(n))
	executed := atomic.Int32{}
	for i := 0; i < epochs * n - 1; i++ {
		go func() {
			b.Wait()
			executed.Add(1)
		}()
	}

	time.Sleep(time.Millisecond * 200)

	if (int(executed.Load()) != (epochs - 1) * n) {
		t.Fatal("Executed too early")
	}

	b.Wait()
	executed.Add(1)

	time.Sleep(time.Millisecond * 200)
	if (int(executed.Load()) != epochs * n) {
		t.Fatal("Didn't execute at all")
	}
}

func TestBarrier_MissedWakeup(t *testing.T) {
	b := NewBarrier(3)

	var wg sync.WaitGroup
	started := make(chan struct{}, 3)

	for i := 0; i < 3; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			started <- struct{}{}
			b.Wait()
		}(i)
	}

	// Подождём, пока все горутины почти дошли
	time.Sleep(10 * time.Millisecond)

	// Проверим, не зависли ли они
	done := make(chan struct{})
	go func() {
		wg.Wait()
		close(done)
	}()

	select {
	case <-done:
		// всё ок
	case <-time.After(2 * time.Second):
		t.Fatal("deadlock: goroutines stuck in barrier.Wait()")
	}
}