package barrier

import (
	"sync/atomic"
)

type Barrier struct {
	maxWaiting int64
	channel chan struct{}
	waiting atomic.Int64
}

func NewBarrier(size int64) *Barrier {
	return &Barrier{maxWaiting: size, channel: make(chan struct{})}
}

func (barrier *Barrier) Wait() {
	newValue := barrier.waiting.Add(1)
	if newValue == barrier.maxWaiting {
		close(barrier.channel)
		return
	}
	<- barrier.channel
}