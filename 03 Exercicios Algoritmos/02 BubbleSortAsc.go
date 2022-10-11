package main

import (
	"fmt"
)

// DEFINE VARIÁVEIS DE CONTROLE
var i, j int

func main() {

	//CRIA O ARRAY COM A QUANTIDADE DE ELEMENTOS, E INPUTA ESSES NÚMEROS
	numeros := [10]int{}
	for i = 0; i < len(numeros); i++ {
		fmt.Scanf("%d", &numeros[i])
	}

	//MOSTRA O ARRAY INICIAL
	fmt.Printf("\n\nO = %d\n", numeros)

	//FAZ A REORGANIZAÇÃO BUBBLE SORT
	for i = 0; i < len(numeros)-1; i++ {
		for j = 0; j < len(numeros)-i-1; j++ {
			if numeros[j] > numeros[j+1] {
				numeros[j], numeros[j+1] = numeros[j+1], numeros[j]
				fmt.Printf("%d = %d\n", i, numeros)
			}
		}
	}
}
