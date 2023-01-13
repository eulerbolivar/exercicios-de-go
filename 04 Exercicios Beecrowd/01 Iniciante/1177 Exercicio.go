package main
 
import (
    "fmt"
)
 
func main() {

    vetor := make([]int, 1000)
	var teste int
	fmt.Scanf("%d", &teste)
	controle := teste

	for i := 0; i < len(vetor); i++{
		vetor[i] = teste - controle
		fmt.Printf("N[%d] = %d\n", i, vetor[i])
		controle--
		if controle == 0{
			controle = teste
		}
	}

}
