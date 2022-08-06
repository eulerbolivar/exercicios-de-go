package main
 
import (
    "fmt"
)

var a int
var b int
var c int
var d int

func main() {
    fmt.Scanf("%v", &a)
    fmt.Scanf("%v", &b)
    fmt.Scanf("%v", &c)
    fmt.Scanf("%v", &d)

    diferenca := (a * b) - (c * d)
    fmt.Printf("DIFERENCA = %v\n", diferenca)

}
