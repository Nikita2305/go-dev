package main 

/*
https://leetcode.com/problems/max-consecutive-ones-iii/description/?envType=study-plan-v2&envId=leetcode-75
*/

import (
	"fmt"
)

func longestOnes(nums []int, k int) int {
	maxSegment := 0
	j := 0
	nullsInSegment := 0
	for i := range nums {
		for j < len(nums) && nullsInSegment <= k {
			if nums[j] == 0 && nullsInSegment == k {
				break
			} else if nums[j] == 0 {
				nullsInSegment++
			}
			j += 1
		}
		maxSegment = max(maxSegment, j - i)
		if nums[i] == 0 {
			nullsInSegment--
		}
	}
	return maxSegment
}

func main() {
	fmt.Println(longestOnes([]int{1,0,0,0,1,1,0,1,1,0,0}, 1))
}