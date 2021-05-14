package main

import (
	"reflect"
	"testing"
)

func TestLongestSubstringWithKRepeatChars(t *testing.T) {
	cases := []struct {
		str string
		k   int
		ans int
	}{
		{
			str: "aaabb",
			k:   3,
			ans: 3,
		}, {
			str: "ababbc",
			k:   2,
			ans: 5,
		}, {
			str: "geeksforgeeks",
			k:   2,
			ans: 2,
		},
	}

	for i, c := range cases {
		res := calcLongestSubstring(c.str, c.k)
		if !reflect.DeepEqual(res, c.ans) {
			t.Errorf(" test case %d failed , actual %#v, expected %#v", i, res, c.ans)
		}
	}
}
