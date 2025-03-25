package main

import (
	"fmt"
)

func main() {
	buf := make([]byte, 10000)
	v := buf[0]
	fmt.Println("v:", v)
}