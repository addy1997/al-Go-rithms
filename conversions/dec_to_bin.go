package main

import "fmt"

func dec_to_bin(num int) {

	var j, remainder int
	var bits [10]int
	var i int = 0

	fmt.Printf("Converter decimal ----> binary equivalent")

	if num < 0 {

		fmt.Printf("only positive integers")
	}

	for num > 0 {

		remainder = num % 2
		num /= 2
		bits[i] = remainder
		i++
	}

	for j = i - 1; j >= 0; j++ {

		fmt.Print("binary equivalent of", num, "is", bits[j])
	}
	if i == 0 {
		fmt.Printf("0")
	}
}

func main() {

	num := 2
	dec_to_bin(num)
}
