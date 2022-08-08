package main

import (
	"fmt"
)

var tempo int
var segundos int
var minutos int
var horas int

var contador int
var i, j, k int

func main() {
	fmt.Scanf("%d", &tempo)

    //CONTROLE DE QUANTAS HORAS
	for i = 0; i < tempo; i++{
        contador++
        if contador == 3600{
            horas++
            contador = 0
        }
    } 

    if (horas > 0){
    tempo = tempo - (3600 * horas)
    }

    //CONTROLE DE QUANTOS MINUTOS
    contador = 0
    for j = 0; j < tempo; j++ {
        contador++
        if contador == 60{
            minutos += 1
            contador = 0
        }
    }

    if minutos > 0{
        tempo = tempo - (60 * minutos)
        }

    //CONTROLE DE QUANTOS SEGUNDOS
    contador = 0
    for k = 0; k < tempo; k++{
        segundos += 1
    }

    fmt.Printf("%d:%d:%d\n", horas, minutos, segundos)
}
