package main

import (
	"fmt"
)

var x int = 2001

func main() {

	for {
		fmt.Printf("%v\n", x)
		x++
		if (x == 2022){
			break
		}
		
	}
}
