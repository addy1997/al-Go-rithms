//program to find perfect digital invariants
package main

import "fmt"

func power(x int, y int) int {

	if y == 0 {

		return 1
	}

	if y%2 == 0 {

		return (power(x, y/2) * power(x, y/2))
	}

	return (x * power(x, y/2) * power(x, y/2))
}

func isPDI(x int) bool {

	var fixedPow int
	var temp int = x
	var sum int = 0
	var rem int = 0
	var fixedNum int = 1
	var r int

	for fixedNum == 1 {

		fixedPow++
		for temp > 0 && rem > 0 {

			rem = temp % 10
			sum += power(r, fixedPow)
			temp /= 10

		}

		if sum == x {
			return true
		}

		if sum > 0 {
			return false
		}
	}
	return true
}

func main() {

	var N int = 4150
	x := 1
	y := 2

	power(x, y)

	if isPDI(N) {
		fmt.Println("Yes")
	} else {

		fmt.Println("No")
	}
}
