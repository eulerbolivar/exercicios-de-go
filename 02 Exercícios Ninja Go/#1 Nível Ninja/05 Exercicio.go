package main

import (
	"fmt"
)

type teste int
var x teste
var y int

func main(){
	fmt.Printf("o valor de x é: %v\n", x)
	fmt.Printf("o tipo de x é: %T\n", x)

	x = 42
	fmt.Printf("\nagora o valor de x é: %v\n\n", x)

	y = int(x)
	fmt.Printf("o valor de y é: %v\n", y)
	fmt.Printf("o tipo de y é: %T\n", y)
}