package main

import "fmt"

// rootVIII remove in place
// 0 ms, faster than 100.00% of Go online submissions

func removeElement(nums []int, val int) int {
	for i := 0; i < len(nums); i++ {
		if nums[i] != val {
			continue
		}
		nums[i] = nums[len(nums)-1]
		nums = nums[:len(nums)-1]
		i--
	}

	return len(nums)
}

func main() {
	fmt.Printf("%v\n", removeElement([]int{3, 2, 2, 3}, 3))
	fmt.Printf("%v\n", removeElement([]int{0, 1, 2, 2, 3, 0, 4, 2}, 2))
}
