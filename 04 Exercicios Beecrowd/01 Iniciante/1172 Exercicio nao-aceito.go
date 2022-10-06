package main
 
import (
    "fmt"
)
 
func main() {
    vetor := make([]int, 10)

    for i := 0; i < 10; i++{
        fmt.Scanf("%d", &vetor[i])
        if vetor[i] <= 0{
            vetor[i] = 1
        }
    }

    for i := 0; i < 10; i++{
        fmt.Printf("X[%d] = %d\n", i, vetor[i])
    }
}
