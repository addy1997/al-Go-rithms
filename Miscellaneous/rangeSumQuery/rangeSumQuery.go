/*
Given an array of n elements. We need to answer q
queries telling the sum of elements in range l to
r in the array. Also the array is not static i.e
the values are changed via some point update query.
*/

package main

import (
	"fmt"
	"math"
)

var block []int
var blk_sz int

func update(arr []int, idx, val int) {
	blk_number := idx / blk_sz
	block[blk_number] = block[blk_number] + val - arr[idx]
	arr[idx] = val
}

func prepocess(arr []int) {
	blk_idx := -1
	blk_sz = int(math.Sqrt(float64(len(arr))))
	block = make([]int, len(arr))
	for i := 0; i < len(arr); i++ {
		if i%blk_sz == 0 {
			blk_idx++
		}
		block[blk_idx] += arr[i]
	}
}

func query(arr []int, l, r int) int {
	sum := 0
	for l < r && l%blk_sz != 0 && l != 0 {
		sum += arr[l]
		l++
	}

	for l+blk_sz <= r {
		sum += block[l/blk_sz]
		l += blk_sz
	}

	for l <= r {
		sum += arr[l]
		l++
	}
	return sum
}

func main() {
	arr := []int{1, 5, 2, 4, 6, 1, 3, 5, 7, 10}
	prepocess(arr)
	sm1 := query(arr, 3, 8)
	sm2 := query(arr, 1, 6)
	fmt.Printf(" value of sum for 3, 8 is %d\n", sm1)
	fmt.Printf(" value of sum for 1, 6 is %d\n", sm2)
	update(arr, 8, 0)
	sm3 := query(arr, 8, 9)
	fmt.Printf(" value of sum for 7, 8 is %d\n", sm3)
}
