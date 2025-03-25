package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	fmt.Println("Num CPU:", runtime.NumCPU())
	fmt.Println("GOMAXPROCS:", runtime.GOMAXPROCS(0))

	for i := 0; i < 100000; i++ {
		go func(i int) {
			time.Sleep(10 * time.Second) // чтобы горутина не завершилась
		}(i)
	}

	fmt.Println("Started 100000 goroutines")
	time.Sleep(15 * time.Second) // ждём, чтобы можно было проанализировать
}
