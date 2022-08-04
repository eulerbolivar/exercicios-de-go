package main

import (
	"fmt"
)

type teste int
var x teste

func main(){
	fmt.Printf("o valor de x é: %v\n", x)
	fmt.Printf("o tipo de x é: %T\n", x)

	x = 42
	fmt.Printf("agora o valor de x é: %v\n", x)
}