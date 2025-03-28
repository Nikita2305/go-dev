package main

import (
	"fmt"
	"sort"
)

func closeStrings(word1 string, word2 string) bool {
    m1 := map[rune]int {}
	for _, v := range word1 {
		m1[v] = m1[v] + 1
	}
	values1 := []int {}
	keys1 := []rune {}
	for k, v := range m1 {
		values1 = append(values1, v)
		keys1 = append(keys1, k)
	}

	m2 := map[rune]int {}
	for _, v := range word2 {
		m2[v] = m2[v] + 1
	}
	values2 := []int {}
	keys2 := []rune {}
	for k, v := range m2 {
		values2 = append(values2, v)
		keys2 = append(keys2, k)
	}

	sort.Ints(values1)
	sort.Ints(values2)

	if len(values1) != len(values2) {
		return false
	}

	for i := range values1 {
		if values1[i] != values2[i] {
			return false
		}
	}

	sort.Slice(keys1, func (i, j int) bool {
		return keys1[i] < keys1[j]
	})
	sort.Slice(keys2, func (i, j int) bool {
		return keys2[i] < keys2[j]
	})

	if len(keys1) != len(keys2) {
		return false
	}

	for i := range keys1 {
		if keys1[i] != keys2[i] {
			return false
		}
	}

	return true
}

func main() {
	fmt.Println(closeStrings("abacaba", "ubucuub"))
}