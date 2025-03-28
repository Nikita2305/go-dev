package main

import (
	"reusable_barrier/barrier"
	"fmt"
	"time"
)

func main() {
	b := barrier.NewBarrier(5)

	for i := 0; i < 10; i++ {
		time.Sleep(time.Millisecond * 200)
		go func() {
			fmt.Println("Start ", time.Now())
			b.Wait()
			fmt.Println("End ", time.Now())
		}()
	}

	time.Sleep(time.Second * 5)
}