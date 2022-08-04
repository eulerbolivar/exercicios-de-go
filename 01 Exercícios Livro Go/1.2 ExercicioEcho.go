package main

import (
	"fmt"
	"os"
)

func main() {
	var s, sep string
	for i := 1; i < len(os.Args); i++{
		s += sep + os.Args[i]
		sep = " "
	}
	
	for i := 1; i < len(os.Args); i++{
		fmt.Printf("o índice é: %v\n", i)
		fmt.Printf("e o valor é: %v\n", os.Args[i])
	}
}