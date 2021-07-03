package main

import "fmt"

func calculate_largest_sum(arr []int) int {
	max_val := 0
	curr_val := 0
	for i := 0; i < len(arr); i++ {
		curr_val = curr_val + arr[i]
		if curr_val < 0 {
			curr_val = 0
		}
		if max_val < curr_val {
			max_val = curr_val
		}
	}
	return max_val
}

func main() {
	arr := []int{-2, -3, 4, -1, -2, 1, 5, -3}
	res := calculate_largest_sum(arr)
	fmt.Println(" Largest sum in arr", res)
}
