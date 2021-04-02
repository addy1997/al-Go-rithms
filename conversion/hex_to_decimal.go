// This program converts all hex numbers except those with 'A','B','C','D','E','F' or 'a','b','c','d','e','f'

package main

import (
	"fmt"
	"math"
)

func convert(hex int) {

	var decimal int = 0
	var remainder int
	var count float64

	for hex > 0 {

		remainder = hex % 10
		decimal += remainder * int(math.Pow(16, count))
		hex /= 10
		count++
	}
	fmt.Println("Decimal Equivalent is-", decimal)
}

func main() {

	hex := 123
	convert(hex)
}
