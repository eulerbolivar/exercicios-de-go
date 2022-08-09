package main

import (
	"fmt"
)

func main() {
	slice := make([]int, 5, 10)

	for i := 0; i < 5; i++{
		slice[i] = i
	}
	
	fmt.Printf("a slice é formada por: %d\n", slice)
	fmt.Printf("a slice tem tamanho de: %d\n", len(slice))
	fmt.Printf("a slice tem capacidade de: %d\n", cap(slice))
}