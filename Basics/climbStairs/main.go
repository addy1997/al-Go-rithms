package main

import "fmt"

/*
LEETCODE: https://leetcode.com/problems/climbing-stairs/

You are climbing a staircase. It takes n steps
to reach the top. Each time you can either climb
1 or 2 steps. In how many distinct ways can you
climb to the top?
*/

func climbStairs(n int) int {
	var items = []int{0, 1, 2}
	for i := 3; i < n+1; i++ {
		items = append(items, items[i-1]+items[i-2])
	}
	return items[n]
}

func main() {
	fmt.Println(climbStairs(45))
	fmt.Println(climbStairs(4))
}
