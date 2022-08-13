package main

import (
	"fmt"
)

var i, j int

func main() {
	numeros := [3]int{}
	for i = 0; i < len(numeros); i++ {
		fmt.Scanf("%d", &numeros[i])
	}

	//DÁ O OUTPUT COM OS 3 NÚMEROS ESCOLHIDOS
	for i = 0; i < len(numeros); i++ {
		fmt.Printf("%d\n", numeros[i])
	}

	fmt.Printf("\n")

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

}
