package main

import (
	"fmt"
	"strings"
)

// FUNÇÃO SOLICITADA PELO EXERCÍCIO
func mostWordsFound(sentences []string) int {

	controle := 0
	for i := 0; i < len(sentences); i++{
		splitSentences := strings.Split(sentences[i], " ")
		if len(splitSentences) > controle{
			controle = len(splitSentences)
		}
	}

	return controle
}

// ENTRADA PADRÃO PARA O USO DA FUNÇÃO
func main(){

	sentences := []string{"alice and bob love leetcode", "i think so too", "this is great thanks very much"}

	resp := mostWordsFound(sentences)
	fmt.Printf("%d\n", resp)
}