package simple_cache

import (
	"sync"
	"sync/atomic"
)

type CounterServiceV3 struct {
	values sync.Map
}

func (c *CounterServiceV3) Inc(key string) {
	actualValue, _ := c.values.LoadOrStore(key, &atomic.Int32{})
	actualValue.(*atomic.Int32).Add(1)
}

func (c *CounterServiceV3) Get(key string) int32 {
	value, ok := c.values.Load(key)
	if !ok {
		return 0
	}
	return value.(*atomic.Int32).Load()
}

func NewCounterServiceV3() *CounterServiceV3 {
	return &CounterServiceV3{values: sync.Map{}}
}