package main

import (
	"fmt"
)

func main(){
	x := make([]string, 26, 26)

	x = []string{"Acre", "Alagoas", "Amapá", "Amazonas",
						"Bahia", "Ceará", "Espírito Santo", "Goiás", "Maranhão", "Mato Grosso", "Mato Grosso do Sul",
						"Minas Gerais", "Pará", "Paraíba", "Paraná", "Pernambuco", "Piauí", 
						"Rio de Janeiro", "Rio Grande do Norte","Rio Grande do Sul",
						"Rondônia", "Roraima", "Santa Catarina", "São Paulo", "Sergipe", "Tocantins"}

	fmt.Printf("\na quantidade de estados é: %d\n", len(x))
	fmt.Printf("a capacidade do slice é: %d\n\n", cap(x))

	for i := 0; i < len(x); i++{
		fmt.Printf("%d = %s\n", i + 1, x[i])
	}
}