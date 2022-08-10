package main

import (
	"fmt"
)

func main() {
	slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	for i, v := range slice {
		fmt.Printf("%d %d\n", i, v)
	}

	fmt.Printf("o tipo do slice é: %T\n\n", slice)
}