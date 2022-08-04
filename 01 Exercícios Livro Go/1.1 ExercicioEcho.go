package main

import (
	"fmt"
	"os"
)

func main(){
	var s, sep string
	for i := 1; i < len(os.Args); i++{
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Printf("O que foi armazenado no S: ")
	fmt.Println(s)

	fmt.Printf("O que foi armazenado no os.Args[0]: ")
	fmt.Println(os.Args[0])
}