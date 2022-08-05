package main

import (
	"fmt"
)

const (
	x = `	esse 
			     texto 
					  está 
						  cru?`)

func main() {
	fmt.Printf("\ndentro do x está escrito: \"%v\"\n\n", x)
}