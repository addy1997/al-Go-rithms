// decimal to octal
package main

import "fmt"

//to convert into toDecimal function
func toOctal(decimal_val int) {

	var octal_val int
	var j int = 1

	for decimal_val != 0 {
		octal_val += (decimal_val % 8) * j
		decimal_val /= 8
		j *= 10
	}
	fmt.Print("the octal equivalent is", octal_val)
}

//driver's function
func main() {
	decimal_val := 10
	toOctal(decimal_val)

}
