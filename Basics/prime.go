package main
import "fmt"

func Isprime(number int){
     
    var i int
    var flag int=0

    for i=2;i<=number/2;i++ {
        if (number%i==0) {
        flag=1
        break
        }
    }
    if (number==1) {
        fmt.Printf("1 is neither prime nor composite.")
    }else{
        if (flag == 0){
            fmt.Printf("%d is a prime number.",number)
        }else{
            fmt.Printf("%d is not a prime number.", number)
        }
    }
}
func main() {
    number:=38
    Isprime(number)
}
