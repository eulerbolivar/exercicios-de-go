package main

import(
	"fmt"
)

func finalValueAfterOperations(operations []string) int {
    var ans int

	for i := 0; i < len(operations); i++{
		if operations[i] == "X++" || operations[i] == "++X"{
			ans += 1
		} 
		if operations[i] == "X--" || operations[i] == "--X"{
			ans -= 1
		}
	}

	return ans
}   

func main(){
	// RECEBE O TEXTO E TRANSFORMA ELE EM UM SLICE DE STRINGS
	entrada := []string{"++X","++X","X++", "--X"}

	resp := finalValueAfterOperations(entrada)
	
	fmt.Printf("%d\n", resp)
}