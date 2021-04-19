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

func bubbleSort(arr []int) {
	isChanged := true
	count := 0
	for isChanged {
		isChanged = false
		for j := 0; j < len(arr)-1-count; j++ {
			if arr[j] > arr[j+1] {
				temp := arr[j+1]
				arr[j+1] = arr[j]
				arr[j] = temp
				isChanged = true
			}
		}
		count++
	}
}

func main() {
	rand.Seed(time.Now().Unix())
	randSlice := rand.Perm(100)
	printSlice(randSlice)
	fmt.Println("\nSorting..")
	bubbleSort(randSlice)
	printSlice(randSlice)
}
