package main

import(
	"fmt"
)

// ENTRADA PADRÃO PARA O USO DA FUNÇÃO
func main(){

	vetor := make([]int, 10)

	fmt.Scanf("%d", &vetor[0])
	for i := 1; i < len(vetor); i++{
		vetor[i] = vetor[i-1] * 2
	}
	
	for i := 0; i < len(vetor); i++{
		fmt.Printf("N[%d] = %d\n", i, vetor[i])
	}
}