package main

import (
	"reflect"
	"testing"
)

func TestWordMatch(t *testing.T) {
	cases := []struct {
		board [][]byte
		word string
		isExist bool
	} {
	{
		board: [][]byte{{'A', 'B', 'C', 'E'}, {'S', 'F', 'C', 'S'}, {'A', 'D', 'E', 'E'}},
		word: "ABCB",
		isExist: false,
	}, {
		board: [][]byte{{'A', 'B', 'C', 'E'}, {'S', 'F', 'C', 'S'}, {'A', 'D', 'E', 'E'}},
		word: "ABCCED",
		isExist: true,
	}, {
		board: [][]byte { {'B', 'N', 'E', 'Y', 'S'}, {'H', 'E', 'D', 'E', 'S'},{'S', 'G', 'N', 'D', 'E'}},
		word: "DES",
		isExist: true,
		},
	}
	for i, c := range cases {
		exist := exist(c.board, c.word)
		if !reflect.DeepEqual(exist, c.isExist) {
			t.Errorf(" test case %d failed , actual %#v, expected %#v", i, exist, c.isExist)
		}
	}
	
}