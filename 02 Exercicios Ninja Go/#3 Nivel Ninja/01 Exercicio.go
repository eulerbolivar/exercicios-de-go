package main

import (
	"fmt"
)

func main(){
	contador := 1

	for i := 1; i <= 500; i++{
		fmt.Printf("\n")
		for j := 1; j <= 20; j++{
			fmt.Printf("%.2v ", contador)
			contador++
		}
	}

	fmt.Printf("\n")
}