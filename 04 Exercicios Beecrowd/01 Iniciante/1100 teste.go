package main
import "fmt"

var i int

func SomaS(vetor[] int) int{
    var somado int
    for i := 0; i < 4; i++{
        somado += vetor[i]
    }

    return somado;
}

func main() {

    numeros := []int{1, 2, 3, 4}

    total := SomaS(numeros)
    
    fmt.Printf("%d\n", total)
    fmt.Printf("%s\n", len(string(total)))

}
