package main

import "fmt"

/*
0ms faster than 100.00% of Go online submissions

LEETCODE: https://leetcode.com/problems/length-of-last-word/

Given a string s consists of some words separated
by spaces, return the length of the last word in
the string. If the last word does not exist, return 0.
A word is a maximal substring consisting of non-space
characters only.
*/

func lengthOfLastWord(s string) int {
	var wordlen int = 0
	var collecting bool = false

	for i := len(s) - 1; i >= 0; i-- {
		if s[i] != ' ' && !collecting {
			collecting = true
		}

		if collecting {
			if s[i] == ' ' {
				break
			} else {
				wordlen++
			}
		}

	}

	return wordlen
}

func main() {
	fmt.Println(lengthOfLastWord("Hellow World"))
	fmt.Println(lengthOfLastWord(" "))
	fmt.Println(lengthOfLastWord("a "))
}
