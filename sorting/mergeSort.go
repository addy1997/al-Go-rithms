package main

import (
	"fmt"
	"math/rand"
	"time"
)

func printSlice(arr []int) {
	for _, num := range arr {
		fmt.Printf("%4d", num)
	}
}

func mergeSort(arr []int) {
	if len(arr) > 1 {
		mid := len(arr) / 2

		left := arr[:mid]

		right := arr[mid:]

		mergeSort(left)

		mergeSort(right)

		leftIndex, rightIndex, arrIndex := 0, 0, 0
		for leftIndex < len(left) && rightIndex < len(right) {
			if left[leftIndex] < right[rightIndex] {
				arr[arrIndex] = left[leftIndex]
				leftIndex++
			} else {
				arr[arrIndex] = right[rightIndex]
				rightIndex++
			}
			arrIndex++
		}
		for leftIndex < len(left) {
			arr[arrIndex] = left[leftIndex]
			leftIndex++
			arrIndex++
		}

		for rightIndex < len(right) {
			arr[arrIndex] = right[rightIndex]
			rightIndex++
			arrIndex++
		}
	}
}

func main() {
	rand.Seed(time.Now().Unix())
	randSlice := rand.Perm(100)
	printSlice(randSlice)
	fmt.Println("\nSorting..")
	mergeSort(randSlice)
	printSlice(randSlice)
}
