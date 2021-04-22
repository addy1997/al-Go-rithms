package main

import (
	"fmt"
)

/*
Given an array that contains both positive and negative integers,
find the product of the maximum product subarray.
*/

func MinOfInt(vars ...int) int {
    min := vars[0]

    for _, i := range vars {
        if min > i {
            min = i
        }
    }

    return min
}

func MaxOfInt(vars ...int) int {
    max := vars[0]

    for _, i := range vars {
        if max < i {
            max = i
        }
    }

    return max
}

func maxSubArray(arr []int) int {
	imin, imax, maxValue:= arr[0], arr[0], arr[0]
	for i:= 0; i < len(arr); i++ {
		if arr[i] < 0 {
			imin, imax = imax, imin
		} else if arr[i] == 0 {
			imin, imax = 1, 1
		}
		imax = MaxOfInt(imax, imax * arr[i])
		imin = MinOfInt(imin, imin * arr[i])
		maxValue = MaxOfInt(imax, imin)
	}
	return maxValue
}

func main() {
	arr := []int { 1, -2, -3, 0, 7, -8, -2 }
	res := maxSubArray(arr)
	fmt.Println("max product is", res)
}