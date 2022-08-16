package main

import (
	"fmt"
)

var Hinicial, Minicial, Hfinal, Mfinal int
var Htempo, Mtempo int

func main(){
	fmt.Scanf("%d %d %d %d", &Hinicial, &Minicial, &Hfinal, &Mfinal)

	if Hinicial == Hfinal && Minicial == Mfinal{
		Hfinal += 24
	}

	if Hfinal < Hinicial{
		Hfinal += 24
	}

	if Mfinal < Minicial{
		Mfinal += 60
		Hfinal -= 1
	}

	Htempo = Hfinal - Hinicial
	Mtempo = Mfinal - Minicial

	fmt.Printf("O JOGO DUROU %d HORA(S) E %d MINUTO(S)\n", Htempo, Mtempo)

}