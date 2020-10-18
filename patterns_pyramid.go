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

func floyd(rows int){

     var i,j int
     var num int=1
     
     for i=1;i<=rows;i++{

         for j=1;j<=i;j++{

         fmt.Printf("%d",num)
         
         num++
         }
     fmt.Printf("\n")
     }
}

func Pascal(rows int){

     var space,i,j int
     var coef int
     
     for i = 0; i < rows; i++ {
         for space = 1; space <= rows - i; space++{
             fmt.Printf("  ")

         for j = 0; j <= i; j++ {
             if (j == 0 || i == 0){
                 coef = 1
             }else{
                 coef = coef * (i - j + 1) / j
             }
         fmt.Printf("%4d", coef)
         }
         }
      }
      fmt.Printf("\n")

     
}

func main() {

     rows:=5
     halfpyramid(rows)
     halfpyramid_nums(rows)
     floyd(rows)
     Pascal(rows)
}
