package main
 
import (
    "fmt"
)
 
func main() {
    var qtd, num, somado int

    fmt.Scanf("%d", &qtd)

    for i := 0; i < qtd; i++{
        fmt.Scanf("%d", &num)

        for j := 1; j < num; j++{
            if num % j == 0{
                somado += j
            }
        }
        if somado == num{
            fmt.Printf("%d eh perfeito\n", num)
        } else{
            fmt.Printf("%d nao eh perfeito\n", num)
        }
        somado = 0
    }
}
