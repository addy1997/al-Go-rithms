package main

import "fmt"

func Palindrome(number int) {

    var temp,remainder int
    var reverse int = 0
    temp = number
    for{
        remainder = number%10
        reverse = reverse*10 + remainder
        number/=10
        
    if(number==0){
      break
    }
  }
    if(temp==reverse){
        fmt.Printf("%d is a Palindrome",temp)
    }else{
        fmt.Printf("%d is not a Palindrome",temp)
    }
}    
func main() {
    number:=1661
    Palindrome(number)
}
