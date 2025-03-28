package main

/*
https://leetcode.com/problems/maximum-number-of-vowels-in-a-substring-of-given-length/?envType=study-plan-v2&envId=leetcode-75
*/

import (
	"fmt"
)

func maxVowels(s string, k int) int {
	runes := []rune(s)
    vowels := map[rune]struct{}{'a': {}, 'e': {}, 'o': {}, 'i': {}, 'u': {}}
	// k <= len(s)
	currentCounter := 0
	maxCounter := 0
	for i, v := range runes {
		
		if _, ok := vowels[v]; ok {
			currentCounter += 1
		}
		if i >= k {
			if _, removeOk := vowels[runes[i - k]]; removeOk {
				currentCounter -= 1
			}
		}

		maxCounter = max(maxCounter, currentCounter)
	}
	return maxCounter
}

func main() {
	fmt.Println(maxVowels("abacaba", 3))
}