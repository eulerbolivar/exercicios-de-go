package main
 
import (
    "fmt"
)
 
var num, fat int

func main() {
    fmt.Scanf("%d", &num)

    qtd := num
    for i := 1; i < qtd; i++{
        num = num * (qtd - i)
    }
    
    fmt.Printf("%d\n", num)
}
