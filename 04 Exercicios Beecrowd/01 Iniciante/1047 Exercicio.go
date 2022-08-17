package main

import (
	"fmt"
)

var Hinicial, Minicial, Hfinal, Mfinal int
var horas, minutos, TempoH, TempoM int

func main(){
	fmt.Scanf("%d %d %d %d", &Hinicial, &Minicial, &Hfinal, &Mfinal)

	//DETERMINA O MÓDULO DO TEMPO
	if Hfinal > Hinicial{
		horas =  Hfinal - Hinicial
	}else{
		horas =  Hinicial - Hfinal	
	}

	if Mfinal > Minicial{
		minutos =  Mfinal - Minicial
	}else{
		minutos =  Minicial - Mfinal	
	}

	//DISTRIBUI E TRANSFORMA OS MINUTOS DO "TEMPO" DIVIDIDOS EM HORAS E MINUTOS
	TempoM = (horas * 60) + minutos
	for i := 0; i < TempoM; i++{
		if i == 60{
			TempoH += 1
			TempoM -= 60
			i = 0
		}
	}

	if Hfinal == Hinicial && Mfinal == Minicial{
		TempoH = 24
	}
	
	fmt.Printf("O JOGO DUROU %d HORA(S) E %d MINUTO(S)\n", TempoH, TempoM)

}