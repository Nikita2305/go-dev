package simple_cache

import (
	"testing"
	"sync"
)

// var BuildCache = NewCounterService
// var BuildCache = NewCounterServiceV2
var BuildCache = NewCounterServiceV3

func TestCounter_Basic(t *testing.T) {
	c := BuildCache()
	c.Inc("foo")
	c.Inc("foo")
	c.Inc("bar")

	if got := c.Get("foo"); got != 2 {
		t.Errorf("foo: expected 2, got %d", got)
	}

	if got := c.Get("bar"); got != 1 {
		t.Errorf("bar: expected 1, got %d", got)
	}
}

func TestCounter_Concurrent(t *testing.T) {
	c := BuildCache()
	var wg sync.WaitGroup

	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := 0; j < 1000; j++ {
				c.Inc("x")
			}
		}()
	}

	wg.Wait()

	if got := c.Get("x"); got != 100_000 {
		t.Errorf("x: expected 100000, got %d", got)
	}
}

func BenchmarkCounter_Inc(b *testing.B) {
	c := BuildCache()

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			c.Inc("hot-key")
		}
	})
}

func BenchmarkCounter_Get(b *testing.B) {
	c := BuildCache()
	c.Inc("read")

	b.ResetTimer()

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			_ = c.Get("read")
		}
	})
}
