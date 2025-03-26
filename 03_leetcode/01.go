package main

import (
	"runtime/debug"
	"fmt"
)

/*
https://leetcode.com/problems/merge-strings-alternately/?envType=study-plan-v2&envId=leetcode-75
*/

// My solution
func mergeAlternately(word1 string, word2 string) string {
    var ret []rune
	slice1 := []rune(word1)
	slice2 := []rune(word2)
	for i := 0; i < max(len(slice1), len(slice2)); i++ {
		if i < len(slice1) {
			ret = append(ret, slice1[i])
		}
		if i < len(slice2) {
			ret = append(ret, slice2[i])
		}
	}
	return string(ret)
}

// Lets consider other solution from leetcode - doesn't work for utf-8
func mergeAlternatelyLeetcode(word1 string, word2 string) string {
    var result string;
    length := len(word1);

    if len(word2) < len(word1){
        length= len(word2);
    }

    for i := 0; i < length; i++{
        result += string(word1[i]);
        result+=string(word2[i]);
    }

    if len(word1) > len(word2){
        result += word1[len(word2):]
    } else if len(word2) > len(word1){
        result += word2[len(word1):]
    }

    return result;
}

// And one more
func mergeAlternatelyLeetcodeV2(word1 string, word2 string) string {
    mergedWord := []rune{};
	for i := 0; i < len(word1) + len(word2); i++ { // I suppose error is here - len(string) is inappropriate when considering runes
        if len(word1) > i {
		    mergedWord = append(mergedWord, rune(word1[i])); // and here also 
        }
        if len(word2) > i {
            mergedWord = append(mergedWord, rune(word2[i]));
        }
	}
	return string(mergedWord)
}

// Now lets try build some simple testing infrastructure

type TestSuite struct {
	lhs string
	rhs string
	expected string
}

var suites = []TestSuite{
	{"abc", "defg123", "adbecfg123"},
	{"фыв", "defg123", "фdыeвfg123"},
} 

func testFunc(testing func (string, string) string) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("\nError while testing, stack:")
			debug.PrintStack()
		}
	}()

	for _, suite := range suites {
		if testing(suite.lhs, suite.rhs) != suite.expected {
			panic("fail!")
		}
	}
}

func main() {
	testFunc(mergeAlternately)
	testFunc(mergeAlternatelyLeetcode)
	testFunc(mergeAlternatelyLeetcodeV2)
}