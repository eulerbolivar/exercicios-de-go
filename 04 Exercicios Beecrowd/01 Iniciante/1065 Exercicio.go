package main

import (
    "fmt"
)

var numeros [5]int
var i, pares int

func main() {
    
    for i = 0; i < 5; i++{
        fmt.Scanf("%d", &numeros[i])    
    }
    
    for i = 0; i < 5; i++{
        if (numeros[i] % 2) == 0{
            pares += 1
        }
    }
    
    fmt.Printf("%d valores pares\n", pares)
}