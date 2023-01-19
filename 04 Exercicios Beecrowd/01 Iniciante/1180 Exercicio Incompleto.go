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
}