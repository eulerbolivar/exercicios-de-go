package main
 
import (
    "fmt"
)
 
func main() {
    var qtd, contador int
    fmt.Scanf("%d", &qtd)

    for i := 0; i < qtd; i++{
        for j := 0; j < 3; j++{
            contador++
            fmt.Printf("%d ", contador)
        }
        contador++
        fmt.Printf("PUM\n")
    }
}
