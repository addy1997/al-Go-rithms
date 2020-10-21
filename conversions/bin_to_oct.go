// binary to octal
package main

import "math"
import "fmt"

func convert(binary_num int) int {

	var octal_num int = 0
	var decimal_num int = 0
	var i float64 = 0.0

	for binary_num != 0 {

		decimal_num += (binary_num % 10) * int(math.Pow(2, i))
		i++
		binary_num /= 10
	}
	i = 1

	for decimal_num != 0 {

		octal_num += (decimal_num % 8) * int(i)
		decimal_num /= 8
		i *= 10
  }
  fmt.Println("the octal equivalent is", octal_num)
return octal_num
}

func main() {

	binary_num := 10000
	convert(binary_num)

}
