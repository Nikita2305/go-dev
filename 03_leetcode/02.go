package main

/*
https://leetcode.com/problems/greatest-common-divisor-of-strings?envType=study-plan-v2&envId=leetcode-75
*/

import (
	"fmt"
)

/*
The first idea that comes to mind - is to use the same logic as we use in gcd.
For given t - answer, we have
str1: |--t--||--t--|
str2: |--t--||--t--||--t--||--t--||--t--||--t--||--t--|
Therefore, we can say, that t is gcdOfStrings(suffix_str2, str1).
As in gcd - we have to decrease suffix_str2 as much as possible.
This algorithm will take about O(n + m), which is not really obvious.
On each step we process about m + n bytes, but in euclid algorithm n,m divide by 2 each few steps.
So after all we will process about n + m + n/2 + m/2 + ... = 2n + 2m = O(m + n)

We cannot improve it as we have to read this strings at least once, therefore lets proceed to implementation

To be clear, lets also consider classic number-gcd - we will prove that the first parameter is divided by 2 each few steps
x, y -> y, z=x%y -> z, w=y%z
examples 
100 99 -> 99, 1 -> 1, 0
100 49 -> 49, 2 -> 2, 1
So if x >= 2y, then it is divided on the first step
If y < x < 2y, then z=x%y=x-y. Is it true, that 2z <= x? 2x-2y <= x? x <= 2y? true.   
*/

func gcdOfStrings(str1 string, str2 string) string {
	if len(str2) == 0 {
		return str1
	}

	times := len(str1) / len(str2)
	for i := 0; i < times * len(str2); i++ {
		if str1[i] != str2[i % len(str2)] {
			return ""
		}
	}

	return gcdOfStrings(str2, str1[times * len(str2):])
}

/*
This solution was correct and shown O(n) complexity
In editorial we had interesting idea i didn't notice.
Strings can have s-gcd of length which is divisor of both len str1 and str2
Thus it is divisor of number-gcd of these lengths. Could this length be less than number gcd?
Actually no, as then it will be divisor of number-gcd.
But we can see that then string-gcd will be also common-string-divisor for str1 and str2.
Thus we can simply find gcd for lengths and check if it really fits both strings.
It could be easily done in O(n + m), though authors suggest checking str1+str2=str2+str1 which does the same.
*/

func main() {
	fmt.Println(gcdOfStrings("1213", "12131213"))
}