package main

import "fmt"

func reverse_factorial(int x) {

	var y int = 1
	var z int = 0

	if x < 0 {

		z=3

	for z = 0; ; y++ {

		if (((x / y) % 1) != 0) || ((x / y) < y) {

			z = 2

		if (x == y) {

			z = 1

    }
		x /= y
  
		if (z == 1) {
			fmt.Print(y + "!")
		}

		if (z == 2) {
			fmt.Print("Not a factorial of a positive whole number")
		}

		if (z == 3) {
			fmt.Print("negative not allowed")
	
  	}

  }  
	}

}

func main(){

	x := 2
	reverse_factorial(x)
  
}
