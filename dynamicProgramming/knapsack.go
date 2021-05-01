#!/usr/bin/env python
# coding: utf-8

# In[ ]:


package main

import "fmt"
import "unsafe"

func max(a, b int) int {
    
    if (a>b){
        return a
    }else{
        return b
    }
}

func knapsack(W int, Wt[] int, v[] int, n int)int{
    
    if (n==0 || W==0){
        return 0
    }
    
    if (Wt[n-1] > W){
        return knapsack(W, Wt, v, n-1)
    }else{
        return max(v[n-1]+ knapsack(W - Wt[n-1], Wt, v, n-1), knapsack(W, Wt, v, n-1))
        
    }
    return 0
}
    
func main(){
        
    var v  =[] int {60,100,120}
    var Wt   =[] int {10,20,30}
    W := 50
    n:= int(unsafe.Sizeof(v) / unsafe.Sizeof(v[0]))
    fmt.Printf("%d", knapsack(W, Wt, v, n))
    
}

