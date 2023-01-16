package main
 
import (
    "fmt"
)
 
func main() {

	var vetor1 := make([]int, 5)
	var vetor2 := make([]int, 5)
	for i := 0; i < 15; i++{
		for j := 0; j < 5; j++{
			fmt.Scanf("%d", &vetor1[i])
		}
	}

}
