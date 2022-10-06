package main
 
import (
    "fmt"
)
 
func main() {
    
    var num, qtd, PrimoAux int

    //PERCORRE OS VALORES PARA IDENTIFICAR OS NÚMEROS PRIMOS
    fmt.Scanf("%d", &qtd)
    for i := 0; i < qtd; i++{
        fmt.Scanf("%d", &num)
        for j := 1; j <= num; j++{
            if num % j == 0{
                PrimoAux++
                //fmt.Printf("o número %d dividido por %d sobrou: %.2f e o PrimoAux vale: %d\n", num, j, float64(num%j), PrimoAux)
            } 
        }
        if PrimoAux == 2{
            fmt.Printf("%d eh primo\n", num)
            PrimoAux = 0
        } else {
            fmt.Printf("%d nao eh primo\n", num)
            PrimoAux = 0
        }  
    }
}
