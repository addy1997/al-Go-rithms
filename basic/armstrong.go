package main

import (
	"fmt"
)

func armstrong(num int) {

	var temp int
	var rem int
	var result int = 0

	temp = num

	for temp > 0 {

		rem = temp % 10
		result += rem * rem*rem
		temp /= 10
	}

	if result == num {

		fmt.Print("Number is an Armstrong number")

	} else {
		fmt.Print("Number is not an Armstrong number")
	}

}

func main() {
	num := 153
	armstrong(num)
}
