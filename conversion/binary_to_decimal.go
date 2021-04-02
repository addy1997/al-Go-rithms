//binary to decimal
package main

import "fmt"

func bin_to_dec(binary_num int) {

	var remainder int
	var decimal_num int = 0
	var temp int = 1

	for binary_num > 0 {

		remainder = binary_num % 10
		binary_num = binary_num / 10
		decimal_num += remainder * temp
		temp *= 2
	}

	fmt.Println("the decimal number is", decimal_num)
}

func main() {

	binary_num := 10000
	bin_to_dec(binary_num)
}
