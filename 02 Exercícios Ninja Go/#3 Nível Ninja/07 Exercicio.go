package main

import (
	"fmt"
)

func main() {
	
	x := 66

	if (x == 5){
		fmt.Printf("O x é igual a 5\n")
	} else if (x <= 5){
		fmt.Printf("O x é igual ou menor que 5\n")
	} else{
		fmt.Printf("O x não é nem 5 e nem menor que 5\n")
	}

}