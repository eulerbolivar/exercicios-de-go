package main
 
import (
    "fmt"
)
 
func main() {

	var j, k int
	var controle = make([]int, 15)
	var vetor1 = make([]int, 5)
	var vetor2 = make([]int, 5)

	for i := 0; i < 15; i++{
		fmt.Scanf("%d", &controle[i])
	}

	// CONTROLE DA QUANTIA MÁXIMA DE ENTRADAS
	for i := 0; i < 15; i++{
		if (controle[i] % 2) == 0{
			if j != 5{
				vetor1[j] = controle[i]
				j++
			} else {
				for i := 0; i < len(vetor1); i++{
					fmt.Printf("par[%d] = %d\n", i, vetor1[i])
				}
				j = 0
			}
		} else { 
			if k != 5{
			vetor2[k] = controle[i]
			k++
		} else {
			for i := 0; i < len(vetor2); i++{
				fmt.Printf("impar[%d] = %d\n", i, vetor2[i])
			}
			k = 0
			}
		}
	}

}
