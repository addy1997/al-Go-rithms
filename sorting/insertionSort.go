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

func insertionSort(arr []int) {
	for i := 1; i < len(arr); i++ {
		if arr[i] < arr[i-1] {
			key := arr[i]
			j := i - 1
			for j >= 0 && arr[j] > key {
				arr[j+1] = arr[j]
				j = j - 1
			}
			arr[j+1] = key
		}

	}
}

func main() {
	rand.Seed(time.Now().Unix())
	randSlice := rand.Perm(100)
	printSlice(randSlice)
	fmt.Println("\nSorting..")
	insertionSort(randSlice)
	printSlice(randSlice)
}
