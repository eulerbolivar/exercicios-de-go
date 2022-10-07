package main

import(
	"fmt"
)


func valorTeste() [][]int {

    var inv int = 2

	// SLICE OF SLICES CRIADO
	teste := make([][]int, 2)
	for i := 0; i < 2; i++{
		teste[i] = make([]int, 3)
	}

	// ARRAY DE CONTROLE PRO EXERCÍCIO
	array := [3]int{1, 2, 3}

	// DISTRIBUIÇÃO DOS NÚMEROS DENTRO DO SLICE OF SLICES
	for i := 0; i < 2; i++{
		for j := 0; j < 3; j++{
			if i == 0{
				teste[i][j] = array[j]
			} else{
				teste[i][j] = array[inv]
				inv--
			}
		}
	}
		
	fmt.Printf("o vetor teste tem: %d\n", teste)
	fmt.Printf("o tipo é %T o tamanho é: %d\n", teste, len(teste))

	ans := teste

	return ans
}

func main(){

	qtd := 2
	tam := 3

	// SLICE DE SLICES CRIADO
	matriz := make([][]int, qtd)
	for i := 0; i < qtd; i++{
		matriz[i] = make([]int, tam)
	}

	test := valorTeste()

	// VALORES ATRIBUÍDOS AOS SLICES
	for i := 0; i < qtd; i++{
		for j := 0; j < tam; j++{
			matriz[i][j] = test[i][j]
		}
	}

	fmt.Printf("%d\n", matriz)

}