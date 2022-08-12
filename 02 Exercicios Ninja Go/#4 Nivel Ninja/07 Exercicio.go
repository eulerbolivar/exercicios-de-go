package main

import (
	"fmt"
)

func main(){
	ss := [][]string{
				[]string{
					"Marcio",
					"Aurelio",
					"Futebol",
				},
				[]string{
					"Fernando",
					"Cruzes",
					"Beisebol",
				},
				[]string{
					"Thiago",
					"Zanatta",
					"Natação",
				},
			}

	for _, v := range ss{
		fmt.Printf("%s\n", v[0])
		for _, itens := range v{
			fmt.Printf("\t %v\n", itens)
		}
	}
}