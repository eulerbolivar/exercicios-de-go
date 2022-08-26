package main

import (
    "fmt"
)

var x int 

func main() {
    fmt.Scanf("%d", &x)
    
    for i := 0; i <= x; i++{
        if((i % 2) != 0){
            fmt.Printf("%d\n", i)
        }
    }
}