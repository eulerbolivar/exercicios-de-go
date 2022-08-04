package main

import (
	"fmt"
)

func main() {
	
	x := 3

	switch{

	case x == 5:
		fmt.Printf("o x é 5\n")
	case x < 5:
		fmt.Printf("o x é menor que 5\n")
		fallthrough
	case x > 5:
		fmt.Printf("o x é maior que 5\n")
	}
}