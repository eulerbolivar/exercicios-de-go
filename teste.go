package main

import (
	"fmt"
)

type hotdog int
var y hotdog

func main() {
	x := 50
	y = 12
	fmt.Printf("\no valor de x é: %v e o tipo de x é: %T\n", x, x)
	fmt.Printf("o valor de y é: %v e o tipo de y é: %T\n\n", y, y)

	y = hotdog(x) //CONVERSÃO DE x PARA O TIPO "HOTDOG"
	fmt.Printf("agora o valor de y passa a ser: %v\n", y)
}