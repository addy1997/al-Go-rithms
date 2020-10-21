package main

import (
	"fmt"
	"math"
)

func convert(octal_num int) int {

	var binary_num int = 0
	var decimal_num int = 0
	var i float64 = 0.0

	for octal_num != 0 {

		decimal_num += (octal_num % 10) * int(math.Pow(8, i))
		i++
		octal_num /= 10
	}
	i = 1

	for decimal_num != 0 {

		binary_num += (decimal_num % 2) * int(i)
		decimal_num /= 2
		i *= 10
	}
	fmt.Println("the binary equivalent is", binary_num)
	return binary_num
}

func main() {

	octal_num := 13
	convert(octal_num)

}
