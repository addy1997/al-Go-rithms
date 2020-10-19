// reference https://www.geeksforgeeks.org/c-program-for-tower-of-hanoi/
package main

import (
	"fmt"
)

func hanoi(num int, from_rod, to_rod, extra_rod byte){

          if (num==1){

          fmt.Print("move disc 1 from rod to rod", from_rod, to_rod)
          }
          
          hanoi(num-1, from_rod, extra_rod, to_rod)

          fmt.Print("move disc %d from rod to rod ", from_rod, to_rod)

          hanoi(num-1, extra_rod, to_rod, from_rod)

}
func main() {
       
     num:=4
     hanoi(num-1, 'a','e','r')

}
