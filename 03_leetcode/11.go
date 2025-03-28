package main

/*
https://leetcode.com/problems/equal-row-and-column-pairs/?envType=study-plan-v2&envId=leetcode-75
*/

import (
	"fmt"
)

func equalPairs(grid [][]int) int {
    entriesRows := map[string]int {}
    for _, v := range grid {
        key := fmt.Sprint(v)
        entriesRows[key] = entriesRows[key] + 1        
    }

	height := len(grid[0])
    trGrid := make([][]int, height)
	for i := 0; i < height; i++ {
		trGrid[i] = make([]int, len(grid))
	}

	for i, v := range grid {
		for j, u := range v {
			trGrid[j][i] = u
		}
	}

    entriesColumns := map[string]int {}
	for _, v := range trGrid {
        key := fmt.Sprint(v)
        entriesColumns[key] = entriesColumns[key] + 1        
    }

	result := 0
	for k, v := range entriesRows {
		result += v * entriesColumns[k]
	}
	return result
}

func main() {
	fmt.Println(equalPairs([][]int{{3,2,1},{1,7,6},{2,7,7}}))
}