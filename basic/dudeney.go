// A Dudeney number is a positive integer that is a perfect cube such that the sum of its decimal digits is equal to the cube root of the number.
// reference: https://en.wikipedia.org/wiki/Dudeney_number

package main
import "fmt"
import "math"

func Dudeney(num int) {

  var rem, temp, cube_root int
  var sum int = 0

  temp = num

  if (temp<0){

    fmt.Println("Please enter a positive number") 
  
  }

  for(temp>0){

    rem = temp%10
    sum+=rem
    temp /= 10
  
  }
  
  cube_root= int(math.Round((math.Pow(float64(num),1.0/3.0))))

  if (cube_root==sum){

    fmt.Printf("it is a Dudeney number")
  }else{
    fmt.Printf("it is not a Dudney number")
  }
}
func main(){

  num:= 33611

  Dudeney(num)

}
