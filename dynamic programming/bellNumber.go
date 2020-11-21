package main
import "fmt"

func bellNumber(n int)int{
    
    s := 100
    
    var bell =  make([s+1][s+1] int)
    
    bell[0][0] = 1
    
    var i, j int
    
    for i=1; i<=n; i++{
        
        bell[i][0] = bell[i-1][i-1]
        
    }
        for j=1; j<=i; j++{
            
            bell[i][j] = bell[i-1][j-1] + bell[i][j-1]
        }
    
    return bell[n][0]
}

func main(){
    
    for n:=0;n<=5;n++{
        
        fmt.Println("Bell number is",bellNumber(n))
    }
}
