package main

import (
	"time"
	"fmt"
	"simple_barrier/barrier"
)

func main() {
	b := barrier.NewBarrier(5)

	for i := 0; i < 5; i++ {
		time.Sleep(time.Millisecond * 200)
		go func() {
			fmt.Println("Waiting")
			b.Wait()
			fmt.Println("Can execute!")
		}()
	}

	time.Sleep(time.Second * 3)

}