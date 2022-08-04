package main

import (
	"fmt"
	//"runtime"
)

const (
	x = 5.4
	y = 10
	z = "texto"
)

func main() {
	fmt.Printf("o valor é: %v e o tipo é: %T\no valor é: %v e o tipo é: %T\no valor é: %v e o tipo é: %T\n\n", x, x, y, y, z, z)
}

// %v = valor
// %T = tipo
// %#U = unicode
// %#x = hexadecimal
