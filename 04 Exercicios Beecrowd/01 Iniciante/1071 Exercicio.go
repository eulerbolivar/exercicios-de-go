package main

import (
    "fmt"
)

var x, y, i, somado int 


func main() {
    fmt.Scanf("%d", &x)
    fmt.Scanf("%d", &y)
    
    maior := x
    menor := y
    
    if y > x{
        maior = y
        menor = x
    }
    
    for i = menor +1; i < maior; i++{
        if (i % 2) != 0{
            somado += i
        }
    }
    
    fmt.Printf("%d\n", somado)
}