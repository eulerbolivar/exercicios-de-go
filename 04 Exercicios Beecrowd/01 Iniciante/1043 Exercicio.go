package main

import (
	"fmt"
)

var i int

func main(){
	numeros := make([]int, 3)

	fmt.Printf("o tamanho é: %d\n", len(numeros))
	fmt.Printf("a capacidade é: %d\n\n", cap(numeros))

	for i = 0; i < 3; i++{
		fmt.Scanf("%d", &numeros[i])
	}

	//DÁ O OUTPUT COM OS 3 NÚMEROS ESCOLHIDOS
	for i = 0; i < 3; i++ {
		fmt.Printf("%d\n", numeros[i])
	}

}