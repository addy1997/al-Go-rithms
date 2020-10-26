package main

import (
	"fmt"
)

func convert(hexadecimal[] byte) {

	var i int = 0

	fmt.Printf("Equivalent binary value: ")

	for hexadecimal[i]>0{

		switch hexadecimal[i] {

		case '0':
			fmt.Printf("0000")
			break

		case '1':
			fmt.Printf("0001")
			break

		case '2':
			fmt.Printf("0010")
			break

		case '3':
			fmt.Printf("0011")
			break

		case '4':
			fmt.Printf("0100")
			break

		case '5':
			fmt.Printf("0101")
			break

		case '6':
			fmt.Printf("0110")
			break

		case '7':
			fmt.Printf("0111")
			break

		case '8':
			fmt.Printf("1000")
			break

		case '9':
			fmt.Printf("1001")
			break

		case 'A':
			fmt.Printf("1010")
			break

		case 'B':
			fmt.Printf("1011")
			break

		case 'C':
			fmt.Printf("1100")
			break

		case 'D':
			fmt.Printf("1101")
			break

		case 'E':
			fmt.Printf("1110")
			break

		case 'F':
			fmt.Printf("1111")
			break

		case 'a':
			fmt.Printf("1010")
			break

		case 'b':
			fmt.Printf("1011")
			break

		case 'c':
			fmt.Printf("1100")
			break

		case 'd':
			fmt.Printf("1101")
			break

		case 'e':
			fmt.Printf("1110")
			break

		case 'f':
			fmt.Printf("1111")
			break

		default:
			fmt.Printf("\n Invalid hexa digit %c ", hexadecimal[i])

		}
		i++

	}

}

func main() {

	hexadecimal := "23abcd32bcde"
	convert(hexadecimal)
}
