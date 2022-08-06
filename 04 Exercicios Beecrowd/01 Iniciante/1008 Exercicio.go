package main
 
import (
    "fmt"
)

var numero int
var horas int
var horaTrabalho float64
var salario float64

func main() {
    fmt.Scanf("%v", &numero)
    fmt.Scanf("%v", &horas)
    fmt.Scanf("%f", &horaTrabalho)

    salario = (float64(horas) * horaTrabalho)

    fmt.Printf("NUMBER = %v\n", numero)
    fmt.Printf("SALARY = U$ %.2f\n", salario)

}
