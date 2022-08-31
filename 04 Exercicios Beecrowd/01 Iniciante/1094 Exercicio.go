package main
import "fmt"

var quantia, coelhos, ratos, sapos, i int

func SomaS(vetor[] int) int{
    var somado int

    for i := 0; i < quantia; i++{
        somado += vetor[i]
    }

    return somado;
}

func main() {

    fmt.Scanf("%d", &quantia)
    
    numeros := make([]int, quantia)
    tipo := make([]int, quantia)
    
    for i = 0; i < quantia; i++{
        fmt.Scanf("%d %c", &numeros[i], &tipo[i])
        if tipo[i] == 'C' || tipo[i] == 'c'{
            coelhos += numeros[i]
        }
        if tipo[i] == 'R' || tipo[i] == 'r'{
            ratos += numeros[i]
        }
        if tipo[i] == 'S' || tipo[i] == 's'{
            sapos += numeros[i]
        }
    }

    total := SomaS(numeros)

    fmt.Printf("Total: %d cobaias\n", total)
    fmt.Printf("Total de coelhos: %d\n", coelhos)
    fmt.Printf("Total de coelhos: %d\n", ratos)
    fmt.Printf("Total de coelhos: %d\n", sapos)

    var PercentC, PercentR, PercentS float64
    

    PercentC = (float64(coelhos) / float64(total)) * 100
    PercentR = (float64(ratos) / float64(total)) * 100
    PercentS = (float64(sapos) / float64(total)) * 100

    fmt.Printf("Percentual de coelhos: %.2f %\n", PercentC)
    fmt.Printf("Percentual de ratos: %.2f %\n", PercentR)
    fmt.Printf("Percentual de sapos: %.2f %\n", PercentS)

}
