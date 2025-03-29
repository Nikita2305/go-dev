package main 

import (
	"fmt"
	"context"
	"time"
	"worker_pool/pool"
)

type Task struct {
	id int
	executeFor time.Duration
}

func (task *Task) Execute(ctx context.Context) error {
	fmt.Println("Starting ", task.id)
	var result error
	select {
		case <- ctx.Done():
			result = fmt.Errorf("Context finished")
		case <- time.After(task.executeFor):
			result = nil
	}
	fmt.Println("Finished", task.id, " with result=", result)
	return result
}

// Пример как можно использовать контекст во внешних библиотеках:

/*
req, err := http.NewRequestWithContext(ctx, http.MethodGet, "https://example.com", nil)
if err != nil {
	log.Fatal(err)
}

resp, err := http.DefaultClient.Do(req)
if err != nil {
	log.Fatal(err) // сюда попадёшь, если контекст истечёт
}
*/

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	pool := pool.NewWorkerPool(ctx, 5)

	pool.Start()

	for i := 0; i < 20; i++ {
		var calculatedDuration time.Duration = time.Duration(1 + (i * 3) % 5) * time.Second
		err := pool.Put(&Task{i, calculatedDuration})
		if err != nil {
			panic(err)
		}
	}
	fmt.Println("Waiting 8 sec")
	time.Sleep(time.Second * 8)
	fmt.Println("Cancelling")

	cancel()

	fmt.Println("Waiting 5 sec")
	time.Sleep(time.Second * 5)
	pool.Shutdown()
}