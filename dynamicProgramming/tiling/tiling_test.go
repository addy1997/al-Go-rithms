package main

import (
	"reflect"
	"testing"
)

func TestSubSet(t *testing.T) {
	cases := []struct {
		n   int
		ans int
	}{
		{
			n:   3,
			ans: 2,
		}, {
			n:   4,
			ans: 3,
		},
	}

	for i, c := range cases {
		res1 := recurseNumOfWays(c.n)
		if !reflect.DeepEqual(res1, c.ans) {
			t.Errorf(" test case %d failed for Recursive solution, actual %#v, expected %#v", i, res1, c.ans)
		}
		res2 := dpNumberOfWays(c.n)
		if !reflect.DeepEqual(res2, c.ans) {
			t.Errorf(" test case %d failed for dp solution, actual %#v, expected %#v", i, res2, c.ans)
		}
	}
}
