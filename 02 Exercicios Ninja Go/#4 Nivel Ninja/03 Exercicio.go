package main

import (
	"fmt"
)

func main() {
	slice := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

	fmt.Printf("do primeiro até o terceiro: %d\n", slice[:3])
	fmt.Printf("do quinto pra frente: %d\n", slice[4:])
	fmt.Printf("do segundo até o sétimo: %d\n", slice[1:7])
	fmt.Printf("do terceiro até o penúltimo: %d\n", slice[3:len(slice)-1])

	fmt.Printf("o tipo do slice é: %T\n\n", slice)
}