package main

import (
	"fmt"
)

func main() {
	
	for i := 65; i < 91; i++{
		fmt.Printf("O Unicode número %v, traz a letra:\n", i)
		for j := 0; j < 3; j++{
			fmt.Printf("%#U\t", i)
		}
		fmt.Printf("\n\n")
	} 
}