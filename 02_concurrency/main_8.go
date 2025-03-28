package main

import (
	"fmt"
	"time"
	"sync"
)

// Assume we don't have the right to copy T
type Task[T any] struct {
	id int
	value *T
}

type Result[R any] struct {
	id int
	value R
}

func ParallelMap[T any, R any](input []T, fn func(T) R) []R {
	tasks := make([]Task[T], len(input))
	for i := range input {
		tasks[i] = Task[T]{id: i, value: &input[i]}
	}

	results := make(chan Result[R], 10)
	var wg sync.WaitGroup

	fmt.Println("start", time.Now())

	for _, task := range tasks {
		wg.Add(1)
		go func(task Task[T]) {
			defer wg.Done()
			results <- Result[R]{task.id, fn(*task.value)}
			fmt.Println("done", task.id, time.Now())
		}(task)
	}

	output := make([]R, len(input))

	go func() {
		wg.Wait()
		close(results)
	}()

	for result := range results {
		output[result.id] = result.value
	}

	return output
}

func WaitAndPrint(n int) int {
	time.Sleep(time.Second * 5)
	fmt.Println(n)
	return n
}

func main() {
	array := []int{1,2,3,1,2,3,1,2,3,1,2,3,1,2,3,1,2,3,1,2,3,1,2,3,1,2,3,1,2,3}
	fmt.Println(ParallelMap(array, WaitAndPrint))
}