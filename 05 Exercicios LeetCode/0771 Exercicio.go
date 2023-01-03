package main

import (
	"fmt"
	"strings"
)

// FUNÇÃO SOLICITADA PELO EXERCÍCIO
func numJewelsInStones(jewels string, stones string) int {

	var ans int
	jewelsSplit := strings.Split(jewels, "")
	stonesSplit := strings.Split(stones, "")

	for i := 0; i < len(jewelsSplit); i++{
		for j := 0; j < len(stonesSplit); j++{
			if jewelsSplit[i] == stonesSplit[j]{
				ans++
			}
		}
	}

	return ans
}


// ENTRADA PADRÃO PARA O USO DA FUNÇÃO
func main(){

	joias := "z"
	pedras := "ZZ"

	resp := numJewelsInStones(joias, pedras)
	fmt.Printf("%d\n", resp)

}