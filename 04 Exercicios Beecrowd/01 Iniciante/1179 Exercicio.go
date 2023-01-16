package main
 
import (
    "fmt"
)
 
func main() {

	var j, k int
	var controle := make([]int, 15)
	var vetor1 := make([]int, 5)
	var vetor2 := make([]int, 5)

	for i := 0; i < 15; i++{
		fmt.Scanf("%d", &controle[i])
	}

	for i := 0; i < 15; i++{
		if (controle[i] % 2) == 0{
			if j != 4{
				vetor1[j]
				fmt.Printf("par[%d] = %d\n", j, vetor1[j])
				j++
			} else {
				j = 0
			}
			
		} else if (controle[i] % 2) != 0{ 
			if k != 4{
			vetor1[j]
			fmt.Printf("par[%d] = %d\n", k, vetor2[k])
			k++
		} else {
			k = 0
			}
		}
	}

}
