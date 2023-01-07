package main

import (
	"fmt"
)

// FUNÇÃO SOLICITADA PELO EXERCÍCIO
func findMinArrowShots(points [][]int) int {
    
	tam := len(points)
	controle := make([][]int, tam)
	c, t, teste, ans := 0, 0, 0, 0

	// DESTRINCHAMENTO DA AMPLITUDE DOS SLICES EM OUTROS SLICES	
	for i := 0; i < tam; i++{
		teste = points[i][1] - points[i][0]
		if teste < 0{
			t = (teste *-1) + 1
		} else{
			t = teste + 1
		}

		// ESTICANDO OS VALORES DENTRO DO SLICE
		controle[i] = make([]int, t)
		c = 0

			for j := points[i][0]; j <= points[i][1]; j++{
				controle[i][c] = j
				c++
			}
			fmt.Printf("%d\n", controle)
		}

	// COMPARANDO SLICES PARA VER QUEM TEM NÚMERO REPETIDO
	for i := 0; i < tam; i++{
		for j := 0; j < len(controle[i]); j++{
			for k := 0; k < tam; k++{
				for l := 0; l < len(controle[k]); l++{
					if i != k{
						if controle[i][j] == controle[k][l]{
							ans += 1
						}
					}
				}
			}
		}
	}

	return ans
}

// ENTRADA PADRÃO PARA O USO DA FUNÇÃO
func main(){

	tam, qtd, c := 2, 4, 0

	// SLICE DE SLICES CRIADO
	matriz := make([][]int, qtd)

	for i := 0; i < qtd; i++{
		matriz[i] = make([]int, tam)
	}

	// INPUT DO DESAFIO
	proposta := []int{10, 16, 2, 8, 1, 6, 7, 12}
	for i := 0; i < qtd; i++{
		for j := 0; j < tam; j++{
			matriz[i][j] = proposta[c]
			c++
		}
	}

	resp := findMinArrowShots(matriz)
	fmt.Printf("%d\n", resp)

}