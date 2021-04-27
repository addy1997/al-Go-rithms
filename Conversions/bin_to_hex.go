package main

import "fmt"

func reverse(numbers [20]int) [20]int {
	for i := 0; i < len(numbers)/2; i++ {
		j := len(numbers) - i - 1
		numbers[i], numbers[j] = numbers[j], numbers[i]
	}
	return numbers
}

func main() {

	var Hex_values = []int{0, 1, 10, 11, 100, 101, 110, 111, 1000, 1001, 1010, 1011, 1100, 1101, 1110, 1111}

	var binary_val int
	var hex [20]int
	var index int = 0
	var i, digit int

	fmt.Printf("Enter binary number:")
	fmt.Scanf("%d", &binary_val)

	for binary_val > 0 {

		digit = binary_val % 10000

		for i = 0; i < 16; i++ {

			if Hex_values[i] == digit {

				if i < 10 {

					hex[index] = (i + 48)
				} else {

					hex[index] = (i - 10) + 65
				}
				index++
				break
			}
		}
		binary_val /= 10000
	}
	hex[index] = '\u0000'
	//reverse(hex)
fmt.Println("The hexidecimal equivalent is:%d", hex[0])
}
