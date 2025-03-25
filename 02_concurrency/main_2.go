package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	runtime.GOMAXPROCS(1)

	go func() {
		for {
			// tight loop — ничего не вызывает, не блокируется
		}
	}()

	time.Sleep(3 * time.Second)
	fmt.Println("Main done")
}
