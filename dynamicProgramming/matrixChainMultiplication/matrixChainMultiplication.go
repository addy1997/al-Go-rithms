package main

import (
	"fmt"
	"math"
)

func multiplicationCost(ms []int) int {
	mSize := len(ms)
	var dp = make([][]int, mSize)
	for i := range dp {
		dp[i] = make([]int, mSize)
	}
	for l := 2; l < mSize; l++ { // length of subchain
		for i := 0; i < mSize-l; i++ { // for each index of matrix
			j := i + l
			dp[i][j] = math.MaxInt64
			for k := i + 1; k < j; k++ { // to find separtor in matrix btw i and j
				tmp := dp[i][k] + dp[k][j] + ms[i]*ms[k]*ms[j]
				if tmp < dp[i][j] {
					dp[i][j] = tmp
				}
			}
		}
	}
	return dp[0][mSize-1]
}

func main() {
	matrixSize := [][]int{{2, 3}, {3, 6}, {6, 4}, {4, 5}}
	sizeArr := make([]int, len(matrixSize)+1)
	idx := 0
	sizeArr[idx] = matrixSize[0][0]
	idx++
	sizeArr[idx] = matrixSize[0][1]
	idx++
	for i := 1; i < len(matrixSize); i++ {
		sizeArr[idx] = matrixSize[i][1]
		idx++
	}
	lowestCost := multiplicationCost(sizeArr)
	fmt.Println("lowest cost is ", lowestCost)
}
