package main

import (
	"fmt"
	"math"
)

var dinheiro, centavos, reais float64
var valor int
var contador int

func main() {

	fmt.Scanf("%f", &dinheiro)
	reais, centavos = math.Modf(dinheiro)
	valor = int(reais)

	fmt.Printf("NOTAS:\n")
	if valor >= 100 {
		contador = valor / 100
		valor = valor % 100
		fmt.Printf("%d nota(s) de R$ 100.00\n", contador)
	} else {
		fmt.Printf("0 nota(s) de R$ 100.00\n")
	}

	if valor >= 50 {
		contador = valor / 50
		valor = valor % 50
		fmt.Printf("%d nota(s) de R$ 50.00\n", contador)
	} else {
		fmt.Printf("0 nota(s) de R$ 50.00\n")
	}

	if valor >= 20 {
		contador = valor / 20
		valor = valor % 20
		fmt.Printf("%d nota(s) de R$ 20.00\n", contador)
	} else {
		fmt.Printf("0 nota(s) de R$ 20.00\n")
	}

	if valor >= 10 {
		contador = valor / 10
		valor = valor % 10
		fmt.Printf("%d nota(s) de R$ 10.00\n", contador)
	} else {
		fmt.Printf("0 nota(s) de R$ 10.00\n")
	}

	if valor >= 5 {
		contador = valor / 5
		valor = valor % 5
		fmt.Printf("%d nota(s) de R$ 5.00\n", contador)
	} else {
		fmt.Printf("0 nota(s) de R$ 5.00\n")
	}

	if valor >= 2 {
		contador = valor / 2
		valor = valor % 2
		fmt.Printf("%d nota(s) de R$ 2.00\n", contador)
	} else {
		fmt.Printf("0 nota(s) de R$ 2.00\n")
	}

	fmt.Printf("MOEDAS:\n")

	if valor == 1 {
		contador = valor / 1
		valor = valor % 1
		fmt.Printf("%d moeda(s) de R$ 1.00\n", contador)
	} else {
		fmt.Printf("0 moeda(s) de R$ 1.00\n")
	}

	centavos += 0.001
	var centavosCont int = int(centavos * 100) 
	var centavos50, centavos25, centavos10, centavos5, centavos1 int
	var i int 
	
	for i = 0; i <= centavosCont; i++{
		if i == 50{
			centavos50++
			centavosCont -= 50
			i = 0
		}
	}

	for i = 0; i <= centavosCont; i++{
		if i == 25{
			centavos25++
			centavosCont -= 25
			i = 0
		}
	}

	for i = 0; i <= centavosCont; i++{
		if i == 10{
			centavos10++
			centavosCont -= 10
			i = 0
		}
	}

	for i = 0; i <= centavosCont; i++{
		if i == 5{
			centavos5++
			centavosCont -= 5
			i = 0
		}
	}

	for i = 0; i <= centavosCont; i++{
		if i == 1{
			centavos1++
			centavosCont -= 1
			i = 0
		}
	}
	fmt.Printf("%d moeda(s) de R$ 0.50\n", centavos50)
	fmt.Printf("%d moeda(s) de R$ 0.25\n", centavos25)
	fmt.Printf("%d moeda(s) de R$ 0.10\n", centavos10)
	fmt.Printf("%d moeda(s) de R$ 0.05\n", centavos5)
	fmt.Printf("%d moeda(s) de R$ 0.01\n", centavos1)
}
