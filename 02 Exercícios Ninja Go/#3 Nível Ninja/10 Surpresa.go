package main

import (
	"fmt"
)

//var x int = 42

func main(){
	for i := 33; i <= 122; i++{
		fmt.Printf("o número decimal: %d, Hexa: %#x, Unicode: %#U\n", i, i, i)
	}
}