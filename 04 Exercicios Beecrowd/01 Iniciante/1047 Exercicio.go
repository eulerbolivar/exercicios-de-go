package main

import (
	"fmt"
)

var Hinicial, Minicial, Hfinal, Mfinal int
var minutos, horas int

func main(){
	fmt.Scanf("%d %d %d %d", &Hinicial, &Minicial, &Hfinal, &Mfinal)

	minutos =  ((Hfinal * 60) - (Hinicial * 60) + (Mfinal - Minicial))

	for i := 0; i < minutos; i++{
		if i == 60{
			horas++
			minutos -= 60
			i = 0
		}
	}

	if Hfinal == Hinicial && Mfinal == Minicial{
		horas = 24
	}
	
	fmt.Printf("O JOGO DUROU %d HORA(S) E %d MINUTO(S)\n", horas, minutos)

}