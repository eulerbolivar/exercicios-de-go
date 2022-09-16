package main
import "fmt"

var qtd int

func main() {
    fmt.Scanf("%d", &qtd)
    
    for i := 1; i < qtd + 1; i++{
        fmt.Printf("%d %d %d\n", i, (i * i), (i * i * i))
    }
}