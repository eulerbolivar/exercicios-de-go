package main

import(
	"fmt"
	"strconv"
)
 /*func minPartitions(n string) int {
   
	ans := n
	resp := int(ans)
	return resp
}*/

func main(){

	var entrada string
	var err error

	//ENVIA O TEXTO E DEFINE A QUANTIDADE DE DIGITOS APOS O SPLIT
	fmt.Scanf("%s", &entrada)

	separados := make([]int, len(entrada))
	texto := []string{entrada}
	
	// SEPARA E RETORNA OS DIGITOS SEPARADOS EM UMA ARRAY 
	for i := 0; i < len(texto); i++{
		separados[i], err = strconv.Atoi(texto[i])
	}

	// VÁLVULA DE ESCAPE PRA POSSÍVEIS ERROS
	if err != nil{
		fmt.Printf("%d\n", err)
	}

	for i := 0; i < len(texto); i++{
		fmt.Printf("%d e %T\n", separados[i], separados)
	}

	/*
	// TRANSFORMA OS DIGITOS SEPARADOS EM INTEIROS
	for i := 0; i < len(texto); i++{
		numeros[i], err = strconv.Atoi(separados[i])
	}

	fmt.Printf("retorno do erro foi: %T\n", err)
	fmt.Printf("o número é: %d e o tipo é: %T\n", numeros, numeros)
	*/
}