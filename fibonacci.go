package main
import "fmt"

func Fibonacci(number int){
     
    var t1 int =0
    var t2 int =1
    var next_term =0
    var i int

    for i=1;i<=number;i++ {
     
        if (i==1) {
            fmt.Print(t1)
            continue
        }
        if (i==2){
            fmt.Print(t2)
            continue
        }
        next_term = t1 + t2
        t1 = t2
        t2 = next_term
        
        fmt.Print(next_term)
    }
}

func main() {
    
    number:=12
    Fibonacci(number)
}
