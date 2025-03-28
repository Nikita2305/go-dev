package main 

/*
https://leetcode.com/problems/removing-stars-from-a-string/description/?envType=study-plan-v2&envId=leetcode-75
*/

import (
	"fmt"
)

func removeStars(s string) string {
    result := make([]rune, 0, len([]rune(s)))
	for _, v := range s {
		if v == '*' {
			if len(result) > 0 {
				result = result[:len(result) - 1]	
			}
		} else {
			result = append(result, v)
		}
	}
	return string(result)
}

func main() {
	fmt.Println(removeStars("leet**cod*e"))
}