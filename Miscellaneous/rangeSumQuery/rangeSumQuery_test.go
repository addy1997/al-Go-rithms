/*
Test program to validate output of range sum query function against number of possible input
*/

package main

import (
	"reflect"
	"testing"
)

func TestRangeSumQuery(t *testing.T) {
	cases := []struct {
		arr     []int
		queries [][]int
		ans     []int
	}{
		{
			arr:     []int{1, 5, 2, 4, 6, 1, 3, 5, 7, 10},
			queries: [][]int{{3, 8}, {1, 6}, {8, 8}},
			ans:     []int{26, 21, 7},
		},
	}
	for _, c := range cases {
		prepocess(c.arr)
		// res := make([]int, len(c.ans))
		for i, q := range c.queries {
			res := query(c.arr, q[0], q[1])
			if !reflect.DeepEqual(res, c.ans[i]) {
				t.Errorf("range sum query for %d, %d for %+v expected %d, result %d \n", q[0], q[1], c.arr, c.ans[i], res)
			}
		}
	}
}

func TestRangeSumQueryUpdate(t *testing.T) {
	arr := []int{1, 5, 2, 4, 6, 1, 3, 5, 7}
	prepocess(arr)
	res := query(arr, 2, 4)
	if res != 12 {
		t.Errorf("range sum query for %d, %d for %+v expected %d, result %d \n", 2, 4, arr, 13, res)
	}
	update(arr, 2, 4)
	res = query(arr, 2, 4)
	if res != 14 {
		t.Errorf("range sum query for %d, %d for %+v expected %d, result %d \n", 2, 4, arr, 14, res)
	}
}
