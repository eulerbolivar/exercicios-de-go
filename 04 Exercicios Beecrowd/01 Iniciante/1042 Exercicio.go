package main

import (
	"fmt"
)
var num1, num2, num3 int
var i, j int

func main() {
	
	//CRIA O ARRAY COM A QUANTIDADE DE ELEMENTOS, E INPUTA ESSES NÚMEROS
	numeros := []int{}
	slice := []int{}
	fmt.Scanf("%d %d %d", &num1, &num2, &num3)

	numeros = append(numeros, num1, num2, num3)
	slice = append(slice, numeros...)

	//FAZ A REORGANIZAÇÃO BUBBLE SORT
	for i = 0; i < len(numeros)-1; i++ {
		for j = 0; j < len(numeros)-i-1; j++ {
			if numeros[j] > numeros[j+1] {
				numeros[j], numeros[j+1] = numeros[j+1], numeros[j]
			}
		}
	}

	//DÁ O OUTPUT COM OS 3 NÚMEROS REORGANIZADOS
	for i = 0; i < len(numeros); i++ {
		fmt.Printf("%d\n", numeros[i])
	}

	fmt.Printf("\n")

	//DÁ O OUTPUT COM OS 3 NÚMEROS ESCOLHIDOS
	for i = 0; i < len(slice); i++ {
		fmt.Printf("%d\n", slice[i])
	}
}
