package main

import (
	"fmt"
)

var x int = 25

func main(){
	
	fmt.Printf("o valor de x é: %v\ndecimal: %d\tbinário: %b\thexadecimal: %#x\n\n", x, x, x, x)
	y := x << 1
	fmt.Printf("o valor de y é: %v\ndecimal: %d\tbinário: %b\thexadecimal: %#x\n\n", y, y, y, y)
}