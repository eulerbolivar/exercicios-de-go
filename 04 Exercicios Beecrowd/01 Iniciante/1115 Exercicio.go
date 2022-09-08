package main

import (
    "fmt"
)

var x, y int

func main() {
    
    
    for {
        fmt.Scanf("%d %d", &x, &y)
        
        if x > 0 && y > 0{
            fmt.Printf("primeiro\n")
        }
        if x < 0 && y > 0{
            fmt.Printf("segundo\n")
        }
        if x < 0 && y < 0{
            fmt.Printf("terceiro\n")
        }
        if x > 0 && y < 0{
            fmt.Printf("quarto\n")
        }
        if x == 0 || y == 0{
            break
        }
        
    }
}