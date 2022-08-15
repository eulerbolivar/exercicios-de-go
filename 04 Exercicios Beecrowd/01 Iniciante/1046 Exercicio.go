package main

import (
	"fmt"
)

var inicial, final int

func main(){
	fmt.Scanf("%d %d", &inicial, &final)

	if final <= inicial{
		final += 24
	}

	tempo := final - inicial

	fmt.Printf("O JOGO DUROU %d HORA(S)\n", tempo)

}