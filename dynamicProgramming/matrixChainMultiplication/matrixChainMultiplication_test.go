package main

import (
	"reflect"
	"testing"
)

func TestMatrixChainMultiplication(t *testing.T) {
	cases := []struct {
		matrixSize [][]int
		cost       int
	}{
		{
			matrixSize: [][]int{{40, 20}, {20, 30}, {30, 10}, {10, 30}},
			cost:       26000,
		}, {
			matrixSize: [][]int{{10, 20}, {20, 30}},
			cost:       6000,
		}, {
			matrixSize: [][]int{{1, 2}, {2, 3}, {3, 4}},
			cost:       18,
		},
	}
	for i, c := range cases {
		sizeArr := make([]int, len(c.matrixSize)+1)
		idx := 0
		sizeArr[idx] = c.matrixSize[0][0]
		idx++
		sizeArr[idx] = c.matrixSize[0][1]
		idx++
		for i := 1; i < len(c.matrixSize); i++ {
			sizeArr[idx] = c.matrixSize[i][1]
			idx++
		}
		cost := multiplicationCost(sizeArr)
		if !reflect.DeepEqual(cost, c.cost) {
			t.Errorf(" test case %d failed , actual %#v, expected %#v", i, cost, c.cost)
		}
	}
}
