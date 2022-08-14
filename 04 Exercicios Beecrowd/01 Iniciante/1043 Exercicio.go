package main

import (
	"fmt"
)

var a, b, c float64

func main(){
	fmt.Scanf("%f %f %f", &a, &b, &c)

	if a + b > c && a + c > b && b + c > a{
		perimetro := a + b + c
		fmt.Printf("Perimetro = %.1f\n", perimetro)
	} else{
		area := ((a + b) * c) / 2
		fmt.Printf("Area = %.1f\n", area)
	}
}