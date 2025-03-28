package main

import (
	"fmt"
	"time"
)

func worker(done <-chan struct{}) {
    for {
        select {
        case <-done:
            fmt.Println("shutting down")
            return
        default:
            fmt.Println("working...")
            time.Sleep(time.Second)
        }
    }
}

func main() {
    done := make(chan struct{})
    go worker(done)

    time.Sleep(3 * time.Second)
    close(done)
    time.Sleep(1 * time.Second)
}
