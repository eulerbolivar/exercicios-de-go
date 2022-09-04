package main

import (
	"fmt"
)

var quantia, coelhos, ratos, sapos int
var total int

func main() {

    fmt.Scanf("%d", &quantia)
    
    var numeros int
    var tipo rune
    
    for i := 0; i < quantia; i++{
        //COLETA A QUANTIDADE E O TIPO DE ANIMAL
        fmt.Scanf("%d", &numeros)
        fmt.Scanf("%c", &tipo)

        //DEBUG
        fmt.Printf("quantidade: %d\n", numeros)
        fmt.Printf("animal: %c\n", tipo)


        //FAZ A DISTRIBUIÇÃO DE TIPOS DE ANIMAIS
        if tipo == 'C' || tipo == 'c'{
            coelhos += numeros
        }
        if tipo == 'R' || tipo == 'r'{
            ratos += numeros
        }
        if tipo == 'S' || tipo == 's'{
            sapos += numeros
        }
        total += numeros
        numeros = 0
    }

    fmt.Printf("Total: %d cobaias\n", total)
    fmt.Printf("Total de coelhos: %d\n", coelhos)
    fmt.Printf("Total de coelhos: %d\n", ratos)
    fmt.Printf("Total de coelhos: %d\n", sapos)

    var PercentC, PercentR, PercentS float64
    

    PercentC = (float64(coelhos) / float64(total)) * 100
    PercentR = (float64(ratos) / float64(total)) * 100
    PercentS = (float64(sapos) / float64(total)) * 100

    fmt.Printf("Percentual de coelhos: %.2f %%\n", PercentC)
    fmt.Printf("Percentual de ratos: %.2f %%\n", PercentR)
    fmt.Printf("Percentual de sapos: %.2f %%\n", PercentS)

}