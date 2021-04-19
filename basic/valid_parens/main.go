package main

import "fmt"

// rootVIII leetcode: 0 ms, faster than 100.00% of Go online submissions

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
