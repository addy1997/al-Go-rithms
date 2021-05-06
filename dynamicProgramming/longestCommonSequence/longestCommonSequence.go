package main

import (
	"fmt"
)

/*
Given two sequences, find the length of longest subsequence present in both of them.
 A subsequence is a sequence that appears in the same relative order, but not necessarily contiguous
*/

func max(first int, rest ...int) (max int) {
	max = first
	for _, value := range rest {
		if value > max {
			max = value
		}
	}
	return
}

func lcs(s1, s2 string) string {
	n := len(s1)
	m := len(s2)
	var dp = make([][]int, n+1)
	// board := make([]string, m*n) creating 2D array with 1D and arithmetic
	// board[i*m + j] = "abc" like board[i][j] = "abc"
	for i := range dp {
		dp[i] = make([]int, m+1)
	}

	for i := 0; i < n+1; i++ {
		for j := 0; j < m+1; j++ {
			if i == 0 || j == 0 {
				dp[i][j] = 0
			} else if s1[i-1] == s2[j-1] {
				dp[i][j] = 1 + dp[i-1][j-1]
			} else {
				dp[i][j] = max(dp[i-1][j], dp[i][j-1])
			}
		}
	}

	res := ""
	k, l := n, m
	for k > 0 && l > 0 {
		if s1[k-1] == s2[l-1] {
			res = res + string(s1[k-1])
			k--
			l--
		} else {
			if dp[k-1][l] > dp[k][l-1] {
				k--
			} else {
				l--
			}
		}
	}
	rns := []rune(res)
	for i, j := 0, len(rns)-1; i < j; i, j = i+1, j-1 {
		rns[i], rns[j] = rns[j], rns[i]
	}
	return string(rns)
}

func main() {
	string1 := "AGGTAB"
	string2 := "GXTXAYB"
	res := lcs(string1, string2)
	fmt.Println("Longest common subsequence", res)
}
