package main
import "fmt"

const MOD int = 1000000007

func Multiply(a[][2] uint,  b[][2] uint){
    
    var x uint  = a[0][0]*b[0][0] + a[0][1]*b[1][0]
    var y uint  = a[0][0]*b[0][1] + a[0][1]*b[1][1]
    var z uint  = a[1][0]*b[0][0] + a[1][1]*b[1][0]
    var w uint  = a[1][0]*b[0][1] + a[1][1]*b[1][1]
    
    a[0][0] = x%uint(MOD)
    a[0][1] = y%uint(MOD)
    a[1][0] = z%uint(MOD)
    a[1][1] = w%uint(MOD)
    
}

func Power(c[][2] uint, n uint){
    
    if(n==1 || n==0){
        
        return 
        
    var F[2][2] uint
        
    F[0][0] = 1
    F[0][1] = 1
    F[1][0] = 1
    F[1][1] = 0
        
    Power(c, n/2)
    Multiply(c,c)
        
    if(n%2==1){
        Multiply(c, F)
    }
    }
}

func main(){
    
    var test int 
    var c[2][2] uint 
    var n uint
    
    fmt.Scanf("%d",&test)
    
    for test=0;;test+=1 {
        
        c[0][0]=1
		    c[0][1]=1
		    c[1][0]=1
		    c[1][1]=0
		
	      fmt.Scanf("%1lu",&n)
	    
	      if(n==0){
	          fmt.Printf("%d\n",0)
	        
	      }else{
	          Power(c,n)
	          fmt.Printf("%1lu\n",c[0][0])
	    }
    }
}

