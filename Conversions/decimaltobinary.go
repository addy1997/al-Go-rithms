package main

import "fmt"
import "unsafe"

func dec_to_bin(num int) int{

  var bits =[][4] int{{0, 0}, {1, 1}, {2, 10}, {3, 11}, {4, 100}}

  var a, b int

  if (num==0){
    num=num%2 * dec_to_bin(num/2)
    fmt.Printf("The value is",num)
  }
  
  size := int(unsafe.Sizeof(bits)/unsafe.Sizeof(bits[0]))

  for i:=0; i<size; i++{
    
    a = dec_to_bin(bits[i][0])

    b = bits[i][1]

    if (a == b){
      fmt.Printf("the quantities",a,"and",b,"are equal")
    }
  }
  return 
}

func main() {

	num := 10
	dec_to_bin(num)
}
