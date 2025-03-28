package simple_cache

import (
	"sync"
)

type CounterService struct {
	values map[string]int
	mutex sync.Mutex
}

func (c *CounterService) Inc(key string) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	c.values[key]++
}

func (c *CounterService) Get(key string) int {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	return c.values[key]
}

func NewCounterService() *CounterService {
	return &CounterService{values: make(map[string]int), mutex: sync.Mutex{}}
}