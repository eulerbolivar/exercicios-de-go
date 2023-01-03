package main

import(
	"fmt"
)

// FUNÇÃO SOLICITADA PELO EXERCÍCIO
func teste()int{
	ans := 10

	return ans
}

// ENTRADA PADRÃO PARA O USO DA FUNÇÃO
func main(){

	resp := teste()
	fmt.Printf("%d\n", resp)

}