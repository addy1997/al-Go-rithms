// this program can only convert 10-->A, 11-->B, 12-->C, 13-->D, 14-->E, 15-->F
package main

import "fmt"

//to convert into toDecimal function
func toHex(decimal_val int) {

	var quotient, remainder int
	var j int = 0
	var hex [10]int

	quotient = decimal_val

	for quotient != 0 {

		remainder = quotient % 16
		if remainder < 10 {
			j += 1
			hex[j] = 48 + remainder
		} else {
			hex[j] = 55 + remainder
			quotient /= 16
		}
	}
	fmt.Println("the hexadecimal equivalent is", hex[j])
}

//driver's function
func main() {
	decimal_val := 11
	toHex(decimal_val)

}
