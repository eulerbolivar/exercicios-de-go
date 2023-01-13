package main
 
import (
    "fmt"
)
 
func main() {

	var x float64
	vetor := make([]float64, 100)

	fmt.Scanf("%f", &x)
	vetor[0] = x

	for i := 1; i < 100; i++{
		vetor[i] = vetor[i - 1]/2
	}
	for i := 0; i < 100; i++{
		fmt.Printf("N[%d] = %.4f\n", i, vetor[i])
	}
}
