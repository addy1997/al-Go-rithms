package main

import (
	"reflect"
	"testing"
)
func TestMaxProductSubArray(t *testing.T) {
	cases := []struct {
		arr []int
		ans int
	} {
	{
		arr: []int{6, -3, -10, 0, 2},
		ans: 180,
	}, {
		arr: []int{2,3,-2,4},
		ans: 6,
	}, {
		arr: []int{-2, -40, 0, -2, -3},
		ans: 80,
		},
	}
	for i, c := range cases {
		res := maxSubArray(c.arr)
		if !reflect.DeepEqual(res, c.ans) {
			t.Errorf(" test case %d failed , actual %#v, expected %#v", i, res, c.ans)
		}
	}
}