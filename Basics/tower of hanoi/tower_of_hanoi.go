// https://www.geeksforgeeks.org/c-program-for-tower-of-hanoi/

package main
import (
	"fmt"
)

func hanoi(num int, from_rod, to_rod, extra_rod byte){

     if (num>1){

        hanoi(num-1, from_rod, extra_rod, to_rod)

        fmt.Println("move disc",num, "from rod",from_rod, "to rod",to_rod)

        hanoi(num-1, extra_rod, to_rod, from_rod)
     } 
}

func main() {

     num:=6
     hanoi(num,'a','e','r')

}
