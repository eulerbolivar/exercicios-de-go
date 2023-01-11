package main
 
import (
    "fmt"
)
 
func main() {

	var vetor = make([]float64, 100)

	for i := 0; i < len(vetor); i++{
		fmt.Scanf("%f", &vetor[i])
		
		if (vetor[i] <= 10){
			fmt.Printf("A[%d] = %.1f\n", i, vetor[i])
		}
	}
    

}
