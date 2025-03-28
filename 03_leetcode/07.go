package main

/*
https://leetcode.com/problems/max-number-of-k-sum-pairs/description/?envType=study-plan-v2&envId=leetcode-75
*/

import (
	"fmt"
)

func maxOperations(nums []int, k int) int {
	entries := make(map[int]int)
	for _, v := range nums {
		entries[v] = entries[v] + 1
	}
	operations := 0
	for key, value := range entries {
		operations += min(value, entries[k - key])
	}
	return operations / 2
}

func main() {
	fmt.Println(maxOperations([]int{1,2,3,4}, 5))
}