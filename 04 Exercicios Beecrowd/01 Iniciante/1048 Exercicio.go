package main

import "fmt"

var Salario, NovoSalario, Reajuste float64
var Porcentagem float64

func main() {
	fmt.Scanf("%f", &Salario)

	switch {
	case Salario > 0 && Salario <= 400:
		Porcentagem = 15.0 / 100
		NovoSalario = Salario + (Salario * Porcentagem)
		Reajuste = NovoSalario - Salario

	case Salario > 400 && Salario <= 800:
		Porcentagem = 12.0 / 100
		NovoSalario = Salario + (Salario * Porcentagem)
		Reajuste = NovoSalario - Salario

	case Salario > 800 && Salario <= 1200:
		Porcentagem = 10.0 / 100
		NovoSalario = Salario + (Salario * Porcentagem)
		Reajuste = NovoSalario - Salario

	case Salario > 1200 && Salario <= 2000:
		Porcentagem = 7.0 / 100
		NovoSalario = Salario + (Salario * Porcentagem)
		Reajuste = NovoSalario - Salario

	case Salario > 2000:
		Porcentagem = 4.0 / 100
		NovoSalario = Salario + (Salario * Porcentagem)
		Reajuste = NovoSalario - Salario
	}

	IntPorcent := int(Porcentagem * 100)

	fmt.Printf("Novo salario: %.2f\n", NovoSalario)
	fmt.Printf("Reajuste ganho: %.2f\n", Reajuste)
	fmt.Printf("Em percentual: %d %%\n", IntPorcent)
}
