package main

import "fmt"

var A, B, C, D int
var contador int

func main() {
    fmt.Scanf("%d %d %d %d", &A, &B, &C, &D)
    
    if B > C{
        contador++
    }
    
    if D > A{
        contador++
    }
    
    if (C + D) > (A + B){
        contador++
    }
    
    if C > 0 && D > 0{
        contador++
    }
    
    if (A % 2 == 0){
        contador++
    }
    
    if contador == 5{
        fmt.Printf("Valores aceitos\n")
    }else{
        fmt.Printf("Valores nao aceitos\n")
    }
}