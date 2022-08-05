package main

import (
	"fmt"
)

func main() {
	var slice[]int
	//EXEMPLO DE ARRAY
	/*
	array := [3]int{1,2,3}
	fmt.Printf("\no valor de array é: %v\n\n", array)
	*/

	fmt.Printf("Escreva os 3 elementos do slice:\n")
	for i := 0; i < 3; i++{
		fmt.Scanf("%v", &slice[i])
	}
	
	fmt.Printf("o valor de slice é: %v\n\n", slice)

	/*
	newSlice := append(slice, 5, 6, 8)
	fmt.Printf("o valor do newSlice é: %v\n\n", newSlice)

	for i := 0; i < len(newSlice); i++{
		fmt.Printf("No índice %v temos o valor: %v\n", i, newSlice[i])
	}
	fmt.Printf("\n")

	fatia := newSlice[:]
	fmt.Printf("A fatia do newSlice nos trouxe: %v\n", fatia)
	*/
}
