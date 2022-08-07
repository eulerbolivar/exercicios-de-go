package main
 
import (
    "fmt"
)

var tempo int
var velocidade int

func main() {
    fmt.Scanf("%d", &tempo)
    fmt.Scanf("%d", &velocidade)

    distancia := velocidade * tempo
    litros := (float64(distancia)/12.0)

    fmt.Printf("%.3f\n", litros)

}
