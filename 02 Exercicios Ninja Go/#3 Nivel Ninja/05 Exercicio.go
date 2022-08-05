package main

import (
	"fmt"
)

func main() {
	for i := 10; i < 100; i++{
		resp := (i % 4)
		fmt.Printf("%v dividido por 4, resta: %v\n", i, resp)
	}
}