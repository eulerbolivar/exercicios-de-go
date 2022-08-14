package main

import (
	"fmt"
)

var a, b, c float64
var i, j, contador int

func main(){
	fmt.Scanf("%f %f %f", &a, &b, &c)

	numeros := []float64{}

	numeros = append(numeros, a, b, c)

	//FAZ A REORGANIZAÇÃO BUBBLE SORT
	for i = 0; i < len(numeros)-1; i++ {
		for j = 0; j < len(numeros)-i-1; j++ {
			if numeros[j+1] > numeros[j] {
				numeros[j+1], numeros[j] = numeros[j], numeros[j+1]
			}
		}
	}

	//DEFINE AS CONDIÇÕES PARA CADA TRIÂNGULO
	if numeros[0] >= numeros[1] + numeros[2]{
	fmt.Printf("NAO FORMA TRIANGULO\n")
	contador++
	}

	if (numeros[0] * numeros[0]) == (numeros[1] * numeros[1]) + (numeros[2] * numeros[2]){
	fmt.Printf("TRIANGULO RETANGULO\n")
	}

	if (numeros[0] * numeros[0]) > (numeros[1] * numeros[1]) + (numeros[2] * numeros[2]) && (contador == 0){
	fmt.Printf("TRIANGULO OBTUSANGULO\n")
	}

	if (numeros[0] * numeros[0]) < (numeros[1] * numeros[1]) + (numeros[2] * numeros[2]){
	fmt.Printf("TRIANGULO ACUTANGULO\n")
	}

	if numeros[0] == numeros[1] && numeros[0] == numeros[2]{
	fmt.Printf("TRIANGULO EQUILATERO\n")
	}

	if (numeros[0] == numeros[1] && numeros[0] != numeros[2]) || (numeros[0] == numeros[2] && numeros[0] != numeros[1]) || (numeros[1] == numeros[2] && numeros[1] != numeros[0]){
	fmt.Printf("TRIANGULO ISOSCELES\n")
	}
}
