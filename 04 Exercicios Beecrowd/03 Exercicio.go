package main
 
import (
    "fmt"
)

var raio float64
var area float64

func main() {
    fmt.Scanf("%v", &raio)

    pi := 3.14159
    raioQuadrado := (raio * raio)
    area := (pi * raioQuadrado)

    fmt.Printf("A=%.4f\n", area)
}
