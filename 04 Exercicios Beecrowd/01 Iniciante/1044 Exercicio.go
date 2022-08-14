package main

import (
	"fmt"
)

var a, b int

func main(){
	fmt.Scanf("%d %d", &a, &b)

	if a % b == 0 || b % a == 0{
		fmt.Printf("Sao Multiplos\n")
	} else{
		fmt.Printf("Nao sao Multiplos\n")
	}

}