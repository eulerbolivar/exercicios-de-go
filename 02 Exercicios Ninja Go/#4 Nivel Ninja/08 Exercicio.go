package main

import (
	"fmt"
)

func main(){
	
	mepe := map[string][]string{
			"Paula": []string{
				"Música", "Café", "Teatro",
			},
			"Roberta": []string{
				"Churrasco", "Cerveja", "Corrida",
			},
	}

	for k, v := range mepe{
		fmt.Printf("%v\n", k)
		for i, h := range v{
			fmt.Printf("%d - %v\n", i, h)
		}
		fmt.Printf("\n")
	}
	
}