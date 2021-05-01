package main

import (
	"reflect"
	"testing"
)

/*
Test program to validate output of getChanges function against number of possible input
*/

func TestCoinChange(t *testing.T) {
	cases := []struct {
		coins []int
		n int
		changes int
	} {
		{
			coins: []int{2, 5, 3, 6},
			n: 10,
			changes: 5,
		},
		{
			coins: []int{1, 2, 3},
			n: 4,
			changes: 4,
		},
		{
			coins: []int{8, 3, 1, 2},
			n: 3,
			changes: 3,
		},
	}
	for i, c := range cases {
		ans := getChanges(c.coins, c.n)
		if !reflect.DeepEqual(ans, c.changes) {
			t.Errorf(" test case %d failed , actual %#v, expected %#v", i, ans, c.changes)
		}
	}
}