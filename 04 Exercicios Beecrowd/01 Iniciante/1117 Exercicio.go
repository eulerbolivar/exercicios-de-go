package main

import (
    "fmt"
)

var nota, somado float64
var contador int
var media float64 = 0

func main() {
    
    for i := 0; i != 2;{
        fmt.Scanf("%f", &nota)
        
        if nota < 0 || nota > 10{
            fmt.Printf("nota invalida\n")
        }else{
            somado += nota
            contador++
            i++
        }
    }
    media = somado / float64(contador)
    
    fmt.Printf("media = %.2f\n", media)
}