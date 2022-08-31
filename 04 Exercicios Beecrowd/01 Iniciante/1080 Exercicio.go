package main
import "fmt"

func main() {
    
    numeros := make([]int, 100)
    
    for i := 0; i < 100; i++{
        fmt.Scanf("%d", &numeros[i])
    }
    
    maior := numeros[0]
    contador := 0
    for i := 0; i < 100; i++{
        if numeros[i] > maior{
            maior = numeros[i]
            contador = i + 1
        }
    }
    fmt.Printf("%d\n", maior)
    fmt.Printf("%d\n", contador)
}
