package main

import (
	"fmt"
)

func main(){

	fmt.Printf("a resposta eh: \n")

	vetor := make([]int, 10)

	for i := 0; i < 10; i++{
		vetor[i] = i
	}

	controle := 10
	fmt.Printf("%d\n", controle)
	for i := 0; i < 10; i++{
		fmt.Printf("[%d] = %d\n", i, vetor[i])
	}
}