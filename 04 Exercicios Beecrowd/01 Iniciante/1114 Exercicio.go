package main

import (
    "fmt"
)

var senha int

func main() {
    
    
    for {
        fmt.Scanf("%d", &senha)
        
        if senha != 2002{
            fmt.Printf("Senha Invalida\n")
        } else{
            fmt.Printf("Acesso Permitido\n")
            break
        }
    }
}