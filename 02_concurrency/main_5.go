package main

import (
	"fmt"
	"time"
)

func worker(id int, jobs <-chan int, results chan<- int) {
    for j := range jobs {
        fmt.Printf("worker %d processing job %d\n", id, j)
        time.Sleep(time.Second) // эмуляция работы
        results <- j * 2
    }
}

func main() {
    jobs := make(chan int, 100)
    results := make(chan int, 100)

    for w := 1; w <= 10; w++ {
        go worker(w, jobs, results)
    }

    for j := 1; j <= 20; j++ {
        jobs <- j
    }
    close(jobs)

    for a := 1; a <= 20; a++ {
        <-results
    }
}
