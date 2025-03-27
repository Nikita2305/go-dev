package main

/*
https://leetcode.com/problems/increasing-triplet-subsequence/description/?envType=study-plan-v2&envId=leetcode-75
*/

import (
	"fmt"
	"math"
)

func increasingTriplet(nums []int) bool {
    mins := make([]int, len(nums) + 1)
	mins[0] = math.MaxInt32
	for i := range nums {
		mins[i + 1] = min(mins[i], nums[i])
	}
	sufMax := math.MinInt32
	for i := len(nums) - 1; i > 0; i-- {
		if nums[i] < sufMax && nums[i] > mins[i] {
			return true
		}
		sufMax = max(sufMax, nums[i])
	}
	return false
}


/*
Неоптимально решили конечно - выделили память, можно без неё
*/


func main() {
	fmt.Println(increasingTriplet([]int {5, 6, 6}))
}