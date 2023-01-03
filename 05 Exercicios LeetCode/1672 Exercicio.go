package main

import (
	"fmt"
)

// FUNÇÃO SOLICITADA PELO EXERCÍCIO
func maximumWealth(accounts [][]int) int {
	maior := 0
	somado := 0

	for i := 0; i < len(accounts); i++{
		for j := 0; j < len(accounts[i]); j++{
			if j == 0{
				somado = 0
			}
			somado += accounts[i][j]
			if somado > maior{
				maior = somado
			}
		}
	} 

	return maior
}

// ENTRADA PADRÃO PARA O USO DA FUNÇÃO
func main(){

	qtd := 3
	tam := 3
	c := 0

	// SLICE DE SLICES CRIADO
	matriz := make([][]int, qtd)
	for i := 0; i < qtd; i++{
		matriz[i] = make([]int, tam)
	}

	// INPUT DO DESAFIO
	proposta := []int{2,8,7,7,1,3,1,9,5}
	for i := 0; i < qtd; i++{
		for j := 0; j < tam; j++{
			matriz[i][j] = proposta[c]
			c++
		}
	}

	resp := maximumWealth(matriz)
	fmt.Printf("%d\n", resp)
}