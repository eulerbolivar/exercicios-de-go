package main

import (
	"fmt"
)

var DDD int

func main(){
	fmt.Scanf("%d", &DDD)

	switch{
	case DDD == 61:
		fmt.Printf("Brasilia\n")
	case DDD == 71:
		fmt.Printf("Salvador\n")
	case DDD == 11:
		fmt.Printf("Sao Paulo\n")
	case DDD == 21:
		fmt.Printf("Rio de Janeiro\n")
	case DDD == 32:
		fmt.Printf("Juiz de Fora\n")
	case DDD == 19:
		fmt.Printf("Campinas\n")
	case DDD == 27:
		fmt.Printf("Vitoria\n")
	case DDD == 31:
		fmt.Printf("Belo Horizonte\n")
	default:
		fmt.Printf("DDD nao cadastrado\n")
	}

}