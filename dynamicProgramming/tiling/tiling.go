/*
Given a "2 x n" board and tiles of size "2 x 1", count the number of ways to tile
 the given board using the 2 x 1 tiles. A tile can either be placed horizontally
 i.e., as a 1 x 2 tile or vertically i.e., as 2 x 1 tile.
*/

package main

import "fmt"

func recurseNumOfWays(n int) int {
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	return recurseNumOfWays(n-1) + recurseNumOfWays(n-2)
}

func dpNumberOfWays(n int) int {
	arr := make([]int, n+1)
	arr[0] = 0
	arr[1] = 1
	for i := 2; i < n+1; i++ {
		arr[i] = arr[i-1] + arr[i-2]
	}
	return arr[n]
}

func main() {
	n := 3
	res := recurseNumOfWays(n)
	fmt.Printf("Number of ways for 4 tiles is %d\n", res)
	res1 := dpNumberOfWays(n)
	fmt.Printf("Number of ways for 4 tiles is %d\n", res1)
}
