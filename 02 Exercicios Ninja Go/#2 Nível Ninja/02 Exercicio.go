package main

import (
	"fmt"
)


func main() {
	x := (5 == 5)
	y := (10 != 6)
	z := (4 <= 4) 
	a := (2 < 6)
	b := (10 >= 12)
	c := (1 > 5)

	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(z)
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)

}