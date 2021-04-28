package main

import "fmt"

/*
finding the total number of ways that an amount of money can be made using specific coins only.
*/

func getChanges(coins []int, n int) int {
	t := make([]int, n+1)
	t[0] = 1
	for _, v := range coins {
		for i := v; i < n+1; i++ {
			t[i] = t[i] + t[i-v]
		}
	}
	return t[n]
}

func main() {
	coins := []int{2, 5, 8}
	n := 10
	res := getChanges(coins, n)
	fmt.Println("number of changes possible is", res)
}