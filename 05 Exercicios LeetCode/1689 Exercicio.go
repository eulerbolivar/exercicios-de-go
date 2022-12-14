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

	// AVALIA A QUANTIDADE DE DECI-BINARIOS SERÃO NECESSÁRIOS
	var ans int

	for i := 0; i < len(texto); i++{
		if separados[i] > ans{
			ans = separados[i]
			}
		}

	return ans
}

func main(){

	var entrada string
	// RECEBE O TEXTO E TRANSFORMA ELE EM UM SLICE DE STRINGS
	fmt.Scanf("%s", &entrada)

	resp := minPartitions(entrada)
	
	fmt.Printf("%d\n", resp)
}