package main

import (
	"fmt"
	"runtime"
	"time"
)

func printMem(tag string) {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	fmt.Printf("[%s] HeapAlloc=%.2fMB Sys=%.2fMB NumGC=%d\n",
		tag, float64(m.HeapAlloc)/1024/1024, float64(m.Sys)/1024/1024, m.NumGC)
}

func testAppend(n int) {
	fmt.Println("== Test append (no preallocation) ==")
	start := time.Now()

	s := make([]int, 0)
	for i := 0; i < n; i++ {
		s = append(s, i)
	}
	duration := time.Since(start)

	printMem("append")
	fmt.Printf("append time: %v\n", duration)
}

func testCopy(n int) {
	fmt.Println("== Test copy (with preallocation) ==")
	start := time.Now()

	s := make([]int, 0, n)
	tmp := make([]int, 1)
	for i := 0; i < n; i++ {
		tmp[0] = i
		s = append(s, tmp...)
	}
	duration := time.Since(start)

	printMem("copy")
	fmt.Printf("copy time: %v\n", duration)
}

func main() {
	n := 10_000_000

	runtime.GC()
	testAppend(n)

	fmt.Println()
	runtime.GC()
	testCopy(n)
}
