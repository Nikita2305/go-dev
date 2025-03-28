package simple_cache

import (
	"sync"
)

type CounterServiceV2 struct {
	values map[string]int
	mutex sync.RWMutex
}

func (c *CounterServiceV2) Inc(key string) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	c.values[key]++
}

func (c *CounterServiceV2) Get(key string) int {
	c.mutex.RLock()
	defer c.mutex.RUnlock()
	return c.values[key]
}

func NewCounterServiceV2() *CounterServiceV2 {
	return &CounterServiceV2{values: make(map[string]int), mutex: sync.RWMutex{}}
}