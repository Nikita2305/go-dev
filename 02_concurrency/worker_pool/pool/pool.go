package pool

import (
	"context"
	"fmt"
	"sync"
	"sync/atomic"
)

type WorkerPoolStatus int

const (
	StatusCreated WorkerPoolStatus = iota
	StatusWorking
	StatusStopped
)

type Task interface {
	Execute(ctx context.Context) error
}

type WorkerPool struct {
	workers int
	tasks chan Task
	ctx context.Context
	wg sync.WaitGroup
	status atomic.Value // WorkerPoolStatus
}

func NewWorkerPool(ctx context.Context, workers int) *WorkerPool {
	tasksQueue := workers * 10
	status := atomic.Value{}
	status.Store(StatusCreated)
	return &WorkerPool{
		workers: workers,
		tasks: make(chan Task, tasksQueue),
		ctx: ctx,
		wg: sync.WaitGroup{},
		status: status,
	}
}

func (pool *WorkerPool) Start() {
	for i := 0; i < pool.workers; i++ {
		pool.wg.Add(1)
		go func() {
			loop:
				for {
					select {
						case <- pool.ctx.Done():
							break loop
						case task, ok := <- pool.tasks:
							if !ok {
								break loop
							}
							err := task.Execute(pool.ctx)
							if err != nil {
								fmt.Println("Error in task, ", err)
							}
					}
				}
			pool.wg.Done()
		}()
	}
	pool.status.Store(StatusWorking)
}

func (pool *WorkerPool) Put(task Task) error {
	if pool.status.Load() != StatusWorking {
		return fmt.Errorf("Pool is not working")
	}
	pool.tasks <- task
	return nil
}

func (pool *WorkerPool) Shutdown() {
	if pool.status.Load() != StatusWorking {
		return
	}
	pool.status.Store(StatusStopped)
	close(pool.tasks)
	pool.wg.Wait()
}