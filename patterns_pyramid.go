package main
import (
	"fmt"
)
func halfpyramid(rows int){
      
     var i, j int
     
     for i=1;i<=rows;i++{
         
        for j=i;j<i;j++{

            fmt.Printf("*")
        }
        fmt.Printf("\n") 
     }
}

func halfpyramid_nums(rows int){
      
     var i, j int
     
     for i=1;i<=rows;i++{
         
        for j=i;j<i;j++{

            fmt.Printf("%d",j)
        }
        fmt.Printf("\n") 
     }
}
func main() {

     rows:=8
     halfpyramid(rows)
     halfpyramid_nums(rows)
}
