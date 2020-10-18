package main 

import "fmt"

func swap(num1 float64, num2 float64){

     num1 = num1 - num2

     num2 = num1 + num2

     num1 = num2 - num1

     fmt.Print("After swapping the first number is", num1)

     fmt.Print("After swapping the second number is", num2)
}

func main(){

     num1 := 66.3
     
     num2 := 44.3

     swap(num1, num2)

}
