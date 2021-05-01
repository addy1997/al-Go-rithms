package main

import (
	"fmt"
)

// rootVIII
// 0 ms, faster than 100.00% of Go online submissions

func plusOne(digits []int) []int {

	var carry bool = false
	for i := len(digits) - 1; i >= 0; i-- {
		if i == len(digits)-1 || carry {
			if digits[i] < 9 {
				digits[i]++
				carry = false
			} else {
				digits[i] = 0
				carry = true
			}
		}
	}
	if carry {
		digits = append([]int{1}, digits...)
	}
	return digits
}

func main() {
	fmt.Printf("%v\n", plusOne([]int{1, 2, 3}))
	fmt.Printf("%v\n", plusOne([]int{9}))
}
