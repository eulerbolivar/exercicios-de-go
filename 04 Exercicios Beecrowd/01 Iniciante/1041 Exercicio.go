package main

import (
	"fmt"
)

var x, y float64

func main(){
	fmt.Scanf("%f %f", &x, &y)

	switch {
	case x == 0 && y == 0:
		fmt.Printf("Origem\n")
	case x == 0 && (y != 0):
		fmt.Printf("Eixo Y\n")
	case y == 0 && (x != 0):
		fmt.Printf("Eixo X\n")
	case x > 0 && y > 0:
		fmt.Printf("Q1\n")
	case x < 0 && y > 0:
		fmt.Printf("Q2\n")
	case x < 0 && y < 0:
		fmt.Printf("Q3\n")
	case x > 0 && y < 0:
		fmt.Printf("Q4\n")
		}
}