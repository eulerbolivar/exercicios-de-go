package main

import (
	"fmt"
)

var x, y, z, i, j int

func main(){
	fmt.Scanf("%d %d %d", &x, &y, &z)

	numeros:= []int{}
	maior:= []int{}
	numeros = append(numeros, x, y, z)

	maior = numeros

	for i = 0; i < len(numeros); i++{
		for j = 0; j < len(numeros); j++{
			if numeros[j] > maior[i]{
				maior[i] = numeros[j]
				maior[i+1] = numeros[j+1]
			}
		}
	}
	for i = 0; i < len(numeros); i++{
		fmt.Printf("%d\n", numeros[i])
	}

	fmt.Printf("\n")

	for i = 0; i < len(maior); i++{
		fmt.Printf("%d\n", maior[i])
	}
	
	

}