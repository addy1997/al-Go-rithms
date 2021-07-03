package main

import (
	"reflect"
	"testing"
)

func TestLargestSumSubArray(t *testing.T) {
	cases := []struct {
		arr []int
		sum int
	}{{
		arr: []int{-2, 1, -3, 4, -1, 2, 1, -5, 4},
		sum: 6,
	}, {
		arr: []int{-2, -3, 4, -1, -2, 1, 5, -3},
		sum: 7,
	},
	}
	for i, c := range cases {
		sum := calculate_largest_sum(c.arr)
		if !reflect.DeepEqual(sum, c.sum) {
			t.Errorf(" test case %d failed , actual %#v, expected %#v", i, sum, c.sum)
		}
	}
}
