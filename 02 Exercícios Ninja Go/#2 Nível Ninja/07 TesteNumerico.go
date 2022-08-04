package main

import (
	"fmt"
	//"runtime"
)
var x int = 100

func main() {
	fmt.Printf("\no valor decimal é: %d\no valor binário é: %b\no valor hexadecimal é: %#x\n\n", x, x, x)
}

// %v = valor
// %T = tipo
// %#U = unicode
// %#x = hexadecimal