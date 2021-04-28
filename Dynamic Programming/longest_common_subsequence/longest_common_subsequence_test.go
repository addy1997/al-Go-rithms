package main

import (
	"reflect"
	"testing"
)

func TestLCS(t *testing.T) {
	cases := []struct {
		s1 string
		s2 string
		ans string
	} {
		{
			s1: "abcdgh",
			s2: "aedfhr",
			ans: "adh",
		}, {
			s1: "aggtab",
			s2: "gxtxayb",
			ans: "gtab",
		}, {
			s1: "abcde",
			s2: "ace",
			ans: "ace",
		},
	}
	for i, c := range cases {
		res := lcs(c.s1, c.s2)
		if !reflect.DeepEqual(res, c.ans) {
			t.Errorf(" test case %d failed , actual %#v, expected %#v", i, res, c.ans)
		}
	}

}