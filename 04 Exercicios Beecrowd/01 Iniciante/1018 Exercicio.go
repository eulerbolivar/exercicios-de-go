package main
 
import (
    "fmt"
)

var valor int
var contador int

func main() {
    
    fmt.Scanf("%d", &valor)
    fmt.Printf("%d\n", valor)
    
    if valor >= 100{
        contador = valor / 100
        valor = valor % 100
        fmt.Printf("%d nota(s) de R$ 100,00\n", contador)
    } else {
        fmt.Printf("0 nota(s) de R$ 100,00\n")
    }

    if valor >= 50{
        contador = valor / 50
        valor = valor % 50
        fmt.Printf("%d nota(s) de R$ 50,00\n", contador)
    } else {
        fmt.Printf("0 nota(s) de R$ 50,00\n")
    }

    if valor >= 20{
        contador = valor / 20
        valor = valor % 20
        fmt.Printf("%d nota(s) de R$ 20,00\n", contador)
    } else {
        fmt.Printf("0 nota(s) de R$ 20,00\n")
    }

    if valor >= 10{
        contador = valor / 10
        valor = valor % 10
        fmt.Printf("%d nota(s) de R$ 10,00\n", contador)
    } else {
        fmt.Printf("0 nota(s) de R$ 10,00\n")
    }

    if valor >= 5{
        contador = valor / 5
        valor = valor % 5
        fmt.Printf("%d nota(s) de R$ 5,00\n", contador)
    } else {
        fmt.Printf("0 nota(s) de R$ 5,00\n")
    }

    if valor >= 2{
        contador = valor / 2
        valor = valor % 2
        fmt.Printf("%d nota(s) de R$ 2,00\n", contador)
    } else {
        fmt.Printf("0 nota(s) de R$ 2,00\n")
    }

    if valor == 1{
        contador = valor / 1
        valor = valor % 1
        fmt.Printf("%d nota(s) de R$ 1,00\n", contador)
    } else {
        fmt.Printf("0 nota(s) de R$ 1,00\n")
    }
}
