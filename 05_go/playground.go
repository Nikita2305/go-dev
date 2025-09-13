package main

import (
	"fmt"
	"unsafe"
)


func ArrayToSlice[T any](a *[5]T) []T {
	return a[:]
}

func main() {
	var arr [5]int
	arr[0] = 2

	slice := ArrayToSlice(&arr)

	array := (*[3]int)(unsafe.Pointer(&slice[0]))
	array[1] = 1
	fmt.Println(arr, slice, array)
}
