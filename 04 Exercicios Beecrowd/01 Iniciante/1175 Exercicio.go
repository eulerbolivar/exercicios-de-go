package main
 
import (
    "fmt"
)
 
func main() {

	vetor := make([]int, 20)
	resp := make([]int, 20)

	for i := 0; i < len(vetor); i++{
		fmt.Scanf("%d", &vetor[i])
	}

	for i := 0; i < (len(resp)/2); i++{
		resp[i] = vetor[len(vetor) - i -1]
		resp[len(vetor) - i -1] = vetor[i]
		
	}

	for i := 0; i < len(vetor); i++{
		fmt.Printf("N[%d] = %d\n", i, resp[i])
	}
}
