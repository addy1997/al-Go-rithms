package main

import (
	"reflect"
	"testing"
)

func TestSubSet(t *testing.T) {
	cases := []struct {
		arr []int
		sum int
		ans bool
	}{
		{
			arr: []int{3, 34, 4, 12, 5, 2},
			sum: 30,
			ans: false,
		}, {
			arr: []int{3, 4, 5, 2},
			sum: 6,
			ans: true,
		}, {
			arr: []int{6, 3, 4},
			sum: 11,
			ans: false,
		},
	}

	for i, c := range cases {
		res1 := recursiveSubsetSum(c.arr, c.sum, len(c.arr)-1)
		if !reflect.DeepEqual(res1, c.ans) {
			t.Errorf(" test case %d failed for Recursive solution, actual %#v, expected %#v", i, res1, c.ans)
		}
		res2 := dpSubSetSum(c.arr, c.sum)
		if !reflect.DeepEqual(res2, c.ans) {
			t.Errorf(" test case %d failed for dp solution, actual %#v, expected %#v", i, res2, c.ans)
		}
	}
}
