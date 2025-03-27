package main

/*
https://leetcode.com/problems/reverse-words-in-a-string/description/?envType=study-plan-v2&envId=leetcode-75
*/

import (
	"fmt"
	"strings"
)

func reverseWords(s string) string {
	var words []string

	for _, s := range strings.Fields(s) {
		if s == "" {
			continue
		}
		words = append(words, s)
	}

	for i := 0; i < len(words) / 2; i++ {
		words[i], words[len(words) - 1 - i] = words[len(words) - 1 - i], words[i]
	}
	
	return strings.Join(words, " ")
}

func main() {
	fmt.Println(reverseWords("  aba caba  daba   "))
}