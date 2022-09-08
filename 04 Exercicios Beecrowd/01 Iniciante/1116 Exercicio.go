package main

import (
    "fmt"
)

var quantia, x, y int

func main() {
    fmt.Scanf("%d", &quantia)
    
    for i := 0; i < quantia; i++{
        fmt.Scanf("%d %d", &x, &y)
        
        dividido := float64(x) / float64(y)
        
        if y != 0{
            fmt.Printf("%.1f\n", dividido)
        }else if y == 0{
            fmt.Printf("divisao impossivel\n")
        }
    }
}