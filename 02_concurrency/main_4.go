package main

import (
	"fmt"
	"time"
	"unsafe"
)

func deep(n int) {
	var buf [1024]byte // ~1 KB на стек
	buf[0] = byte(n)

	var sum byte
	for _, b := range buf {
		sum += b
	}
	fmt.Printf("sum: %d\n", sum)

	printStackPointer(n)

	if n == 0 {
		return
	}
	deep(n - 1)
}

func printStackPointer(n int) {
	var marker int
	fmt.Printf("depth: %4d | SP ~= 0x%x\n", n, uintptr(unsafe.Pointer(&marker)))
}

func main() {
	go deep(200)
	time.Sleep(2 * time.Second)
}
