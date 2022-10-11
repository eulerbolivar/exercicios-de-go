package main

import (
	"fmt"
	"strings"
)

func main(){
	slice := "abcde123456789"

	splitado := strings.Split(slice, "")
	fmt.Printf("o texto é: %s\ne o tipo é: %T\no tamanho é: %d\n", splitado, splitado, len(splitado))
}