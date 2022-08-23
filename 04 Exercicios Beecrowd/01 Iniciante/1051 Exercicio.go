package main

import (
	"fmt"
)

var renda float64

func main(){
	fmt.Scanf("%f", &renda)

	for i := 0; i < int(renda); i++{
		fmt.Printf("o i vale: %d\n", i)
	}

}