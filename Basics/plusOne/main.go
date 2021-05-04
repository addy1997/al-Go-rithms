package main

import (
	"fmt"
)

/*
rootVIII
0 ms, faster than 100.00% of Go online submissions

LEETCODE: https://leetcode.com/problems/plus-one/

Given a non-empty array of decimal digits representing
a non-negative integer, increment one to the integer. The
digits are stored such that the most significant digit is
at the head of the list, and each element in the array
contains a single digit. You may assume the integer does
not contain any leading zero, except the number 0 itself.
*/

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
