package main

import (
	"fmt"
	"strings"
)

// FUNÇÃO SOLICITADA PELO EXERCÍCIO
func minimumSum(num int) int {
    
	controle := make([]int, 2)
	menor := 0
	
	//SPLIT DOS 4 NÚMEROS POSSÍVEIS QUE VIRÃO	
	novoNum := strings.Split(num)

	for i := 0; i < 1; i++{
		menor++
		controle[0] = 1
	}

	ans := 10

	return ans
}

// ENTRADA PADRÃO PARA O USO DA FUNÇÃO
func main(){

	numero := 4009

	resp := minimumSum(numero)

	fmt.Printf("%d\n", resp)

}