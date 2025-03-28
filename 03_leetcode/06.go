package main

/*
https://leetcode.com/problems/container-with-most-water/description/?envType=study-plan-v2&envId=leetcode-75
*/

import (
	"fmt"
	"sort"
	"math"
)

/*
| | | | | | | |
      ^

The first idea that comes to mind is to consider each element as the lowest element in the pair of values
So if we do it, we should find the most distanced element that is more or equal than considered element.
We can do it iterating from right to left, with dekart tree of elements on the suffix and counting max_elem_id on the subtrees.
Then we will have to iterate from left to right also with the same approach. It will be O(n logn) afterall.

But i suppose it is too much for this task, lets try to find simplier approach.
I had an interesting idea to change order of iteration and do it not increasing indices, but increasing values (decreasing actually)
Lets sort values by descending order and then iterate over them.
Then we will increase context of huge values and then consider new "smaller" element of the looked pair.
To find the best fit for our considered x_i, we will have to find x_j > x_i with j = min/max {set_0, ..., set_q}.
Thus we can count max and min for considered indices. 
*/

type Pair struct {
	Value int
	Index int
}

func maxArea1(height []int) int {
    elements := make([]Pair, len(height))
	for i, v := range height {
		elements[i] = Pair{v, i}
	}
	sort.Slice(elements, func (i, j int) bool {
		return elements[i].Value > elements[j].Value
	})
	mx, mn := math.MinInt32, math.MaxInt32 
	var answer int64 = 0
	for _, v := range elements {
		answer = max(answer, int64(v.Value) * int64(v.Index - mn))
		answer = max(answer, int64(v.Value) * int64(mx - v.Index))
		mx = max(v.Index, mx)
		mn = min(v.Index, mn)
	}
	return int(answer)
}

/*
This solution works, but it is still O(n logn) because of sort.
Lets come up with some O(n) two pointers solution (as this chapter tells to:) // didn't notice this hint at first

The new idea is to find the stupides ever solution and then find some improvement.
The stupidest - we do two nested loops and choose the best l and r.
Now lets do some observations. If we ever find some v_l, and then consider l < j, and v[j] < v[l] - we can skip this j.
Then we should consider only j, so that j_1 < j_2 < j3 < ... and v[j_1] < v[j_2] < v[j_3] ..., these j are candidates to be the best L
Also we have the same considerations for candidates of R: v[k_1] < v[k_2] < v[k_3], where k_3 < k_2 < k_1
Then we can use two pointers for L and R candidates (kind of keeping v[l] > v[r])
*/

func maxArea(height []int) int {
	j := len(height) - 1
	answer := 0

	for i := range height {
		if i >= j {
			break
		}

		answer = max(answer, min(height[i], height[j]) * (j - i))
		for height[i] > height[j] && j - 1 > i {
			j -= 1
			answer = max(answer, min(height[i], height[j]) * (j - i))
		}		
	}

	return answer
}

func main() {
	fmt.Println("Answer", maxArea([]int{3,6,8,7,2,1}))
}