/*
Given a set of non-negative integers, and a value sum,
determine if there is a subset of the given set with sum equal to given sum.
*/

package main

import "fmt"

// This function recursively check if sum can be formed from the elements of array in subset
func recursiveSubsetSum(arr []int, sum int, n int) bool {
	if sum == 0 {
		return true
	}

	if n == -1 {
		return false
	}

	if arr[n] > sum {
		recursiveSubsetSum(arr, sum, n-1)
	}

	return recursiveSubsetSum(arr, sum, n-1) || recursiveSubsetSum(arr, sum-arr[n], n-1)
}

// This function check using dp if sum can be formed from the elements of array in subset
func dpSubSetSum(arr []int, sum int) bool {
	dp := make([][]bool, len(arr)+1)
	for i := range dp {
		dp[i] = make([]bool, sum+1)
	}
	// If sum is 0, then answer is true for all arr value
	for i := range dp {
		dp[i][0] = true
	}
	for i := 1; i < sum+1; i++ {
		dp[0][i] = false
	}
	for i := 1; i < len(arr)+1; i++ {
		for j := 1; j < sum+1; j++ {
			if j < arr[i-1] {
				dp[i][j] = dp[i-1][j]
			} else {
				dp[i][j] = dp[i-1][j] || dp[i-1][j-arr[i-1]]
			}
		}
	}
	return dp[len(arr)][sum]
}

func main() {
	arr := []int{3, 34, 4, 12, 5, 2}
	sum := 9
	res1 := recursiveSubsetSum(arr, sum, len(arr)-1)
	res2 := dpSubSetSum(arr, sum)
	fmt.Printf("sum %d can be formed with subarray is %v, %v", sum, res1, res2)
}
