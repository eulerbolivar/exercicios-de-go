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

	for i := 0; i < 15; i++{
		if (controle[i] % 2) == 0{
			if j != 5{
				vetor1[j] = controle[i]
				fmt.Printf("par[%d] = %d\n", j, vetor1[j])
				j++
			} else {
				j = 0
			}
		} else { 
			if k != 5{
			vetor2[k] = controle[i]
			fmt.Printf("impar[%d] = %d\n", k, vetor2[k])
			k++
		} else {
			k = 0
			}
		}
	}

}
