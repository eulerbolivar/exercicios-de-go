package main
 
import (
    "fmt"
)

var peca1 int
var peca2 int

var num1 int
var num2 int

var valor1 float64
var valor2 float64


func main() {
    fmt.Scanf("%v %v %f", &peca1, &num1, &valor1)
    fmt.Scanf("%v %v %f", &peca2, &num2, &valor2)

    custo1 := float64(num1) * valor1
    custo2 := float64(num2) * valor2

    valorTotal := custo1 + custo2
    fmt.Printf("VALOR A PAGAR: R$ %.2f\n", valorTotal)
}
