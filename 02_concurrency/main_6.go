package main

import (
	"fmt"
	"time"
)

func slow() chan int {
    ch := make(chan int)
    go func() {
        time.Sleep(3 * time.Second)
        ch <- 42
    }()
    return ch
}

func main() {
    select {
    case res := <-slow():
        fmt.Println("got result:", res)
    case <-time.After(1 * time.Second):
        fmt.Println("timeout")
    }
}
