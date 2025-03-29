package main

import (
	"fmt"
	"time"
)

func main() {
	var ready bool

	go func() {
		ready = true
	}()
	
	if ready {
		fmt.Println("ready!")
	}

	time.Sleep(100 * time.Millisecond)
}