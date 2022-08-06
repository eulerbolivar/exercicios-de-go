package main
 
import (
    "fmt"
)

var nome string
var salarioFixo float64
var vendas float64
var extra float64
var comissao float64

func main() {
    fmt.Scanf("%s", &nome)
    fmt.Scanf("%f", &salarioFixo)
    fmt.Scanf("%f", &vendas)

    extra = vendas * (15.0/100.0)
    somado := salarioFixo + extra

    fmt.Printf("TOTAL = R$ %.2f\n", somado)
}
