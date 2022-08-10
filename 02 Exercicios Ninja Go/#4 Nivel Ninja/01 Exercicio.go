package main

import (
	"fmt"
)

func main() {
	array := [5]int{}

	for i := 0; i < cap(array); i++{
		array[i] = i
	}

	for i, v := range array {
		fmt.Printf("%d %d\n", i, v)
	}

	fmt.Printf("o tipo do array é: %T\n\n", array)
}