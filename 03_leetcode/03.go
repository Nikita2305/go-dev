package main

/*
https://leetcode.com/problems/kids-with-the-greatest-number-of-candies/description/?envType=study-plan-v2&envId=leetcode-75
*/

import (
	"fmt"
)

func kidsWithCandies(candies []int, extraCandies int) []bool {
    var max_candies = 0 
	for _, v := range candies {
		max_candies = max(max_candies, v)
	}

	results := make([]bool, len(candies))
	for i, v := range candies {
		results[i] = (v + extraCandies >= max_candies)
	}

	return results
}

func main() {
	fmt.Println(kidsWithCandies([]int {1, 3, 5}, 3))
}