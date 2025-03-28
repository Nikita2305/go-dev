package barrier

import (
	"testing"
	"sync/atomic"
	"time"
)

func TestSimple(t *testing.T) {
	n := 5
	b := NewBarrier(int64(n))
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