import (
	"fmt"
	"math"
)

func convert(octal int) int {

	var decimal int = 0
	var i float64

	for octal != 0 {

		decimal += (octal % 10) * int(math.Pow(8, i))
		i += 1.0
		octal /= 10
	}
	fmt.Println("The decimal equivalent is", decimal)
	return decimal
}

func main() {

	octal := 20
	convert(octal)
}
