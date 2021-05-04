package main

import "fmt"

/*
0 ms, faster than 100.00% of Go online submissions

LEETCODE: https://leetcode.com/problems/valid-parentheses/

Given a string s containing just the characters
'(', ')', '{', '}', '[' and ']', determine if the
input string is valid.

An input string is valid if:
Open brackets must be closed by the same type of brackets.
Open brackets must be closed in the correct order.
*/

func isValid(s string) bool {

	var q []rune
	table := map[rune]rune{')': '(', ']': '[', '}': '{'}
	var open uint

	for _, bracket := range s {
		_, ok := table[bracket]
		if !ok {
			q = append(q, bracket)
			open++
			continue
		}
		if open == 0 {
			return false
		}
		end := q[len(q)-1]
		q = q[:len(q)-1]
		open--
		if table[bracket] != end {
			return false
		}
	}

	if open != 0 {
		return false
	}
	return true
}

func main() {
	fmt.Printf("%v\n", isValid("()"))
	fmt.Printf("%v\n", isValid("()[]{}"))
	fmt.Printf("%v\n", isValid("([)]"))
	fmt.Printf("%v\n", isValid("["))
	fmt.Printf("%v\n", isValid("(("))
	fmt.Printf("%v\n", isValid("({{{{}}}))"))
	fmt.Printf("%v\n", isValid("{[]}"))
	fmt.Printf("%v\n", isValid("{[]"))
}
