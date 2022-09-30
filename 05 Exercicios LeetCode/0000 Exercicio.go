package main

import(
	"fmt"
	"strconv"
	"strings"
)
func minPartitions(n string) int {
	var err error

	texto := strings.Split(n, "")
	separados := make([]int, len(texto))
	
	// SEPARA E RETORNA OS DIGITOS SEPARADOS EM UMA ARRAY 
	for i := 0; i < len(texto); i++{
		separados[i], err = strconv.Atoi(texto[i])
	}

	// VÁLVULA DE ESCAPE PRA POSSÍVEIS ERROS
	if err != nil{
		fmt.Printf("%d\n", err)
	}

	// EXPÕE TODOS OS DADOS ARMAZENADOS
	fmt.Printf("aqui será mostrado digito por digito do slice\n")
	for i := 0; i < len(texto); i++{
		fmt.Printf("%d\n", separados[i])
	}

	fmt.Printf("o número é: %d e o tipo é: %T\n", separados, separados)
	
	ans := 5
	return ans
}

func main(){

	var entrada string
	// RECEBE O TEXTO E TRANSFORMA ELE EM UM SLICE DE STRINGS
	fmt.Scanf("%s", &entrada)

	resp := minPartitions(entrada)
	
	fmt.Printf("%d\n", resp)
}