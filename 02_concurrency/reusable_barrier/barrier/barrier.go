package barrier

import (
	"sync"
	"sync/atomic"
)

type Barrier struct {
	size int32
	mx sync.Mutex
	cv *sync.Cond
	currentTicket atomic.Int32
	currentEpoch atomic.Int32 
}

func NewBarrier(size int32) *Barrier {
	barrier := Barrier{
		size: size,
	}

	barrier.cv = sync.NewCond(&barrier.mx)

	return &barrier
}

func (barrier *Barrier) Wait() {
	ticket := barrier.currentTicket.Add(1)
	barrier.mx.Lock()
	defer barrier.mx.Unlock()

	for {
		if ticket == (barrier.currentEpoch.Load() + 1) * barrier.size {
			barrier.currentEpoch.Add(1)
			barrier.cv.Broadcast()
			break
		}
		if ticket < barrier.currentEpoch.Load() * barrier.size {
			break
		}

		barrier.cv.Wait()
	}
}