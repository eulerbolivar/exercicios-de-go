package main

import (
    "fmt"
)

var numeros [5]int
var i, pares, impares, positivos, negativos int

func main() {
    
    for i = 0; i < 5; i++{
        fmt.Scanf("%d", &numeros[i])    
    }
    
    for i = 0; i < 5; i++{
        if (numeros[i] % 2) == 0{
            pares += 1
        } else{
            impares += 1
        }
    
        if (numeros[i] > 0){
            positivos += 1
        } else if (numeros[i] < 0){
            negativos += 1
        }
    }
    
    fmt.Printf("%d valor(es) par(es)\n", pares)
    fmt.Printf("%d valor(es) impar(es)\n", impares)
    fmt.Printf("%d valor(es) positivo(s)\n", positivos)
    fmt.Printf("%d valor(es) negativo(s)\n", negativos)
}