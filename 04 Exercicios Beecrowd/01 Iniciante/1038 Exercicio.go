package main


import "fmt"

var codigo int
var preco, quantidade float64

func main() {
    fmt.Scanf("%d %f", &codigo, &quantidade)
    
    if codigo == 1{
        preco = 4.00 * quantidade
    } else if codigo == 2{
        preco = 4.50 * quantidade
    } else if codigo == 3{
        preco = 5.00 * quantidade
    } else if codigo == 4{
        preco = 2.00 * quantidade
    } else if codigo == 5 {
        preco = 1.50 * quantidade
    }
    fmt.Printf("Total: R$ %.2f\n", preco)
}