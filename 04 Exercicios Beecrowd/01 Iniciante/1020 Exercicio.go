package main
 
import (
    "fmt"
)

var diasVida int
var anos int
var meses int
var dias int

var contador, i int

func main() {
    fmt.Scanf("%d", &diasVida)

    //CONTA OS ANOS DE VIDA
    for i = 0; i < diasVida; i++{
        contador++
        if contador == 365{
            anos++
            contador = 0
        }
    }

    //CONTA OS MESES DE VIDA
    diasVida = contador
    contador = 0
    for i = 0; i < diasVida; i++{
        contador++
        if contador == 30{
            meses++
            contador = 0
        }
    }

    //CONTA OS DIAS DE VIDA
    diasVida = contador
    contador = 0
    for i = 0; i < diasVida; i++{
        dias++
    }

    fmt.Printf("%d ano(s)\n", anos)
    fmt.Printf("%d mes(es)\n", meses)
    fmt.Printf("%d dia(s)\n", dias)
}
