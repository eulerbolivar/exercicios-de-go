package main

import (
	"fmt"
)

func main() {

	switch EsporteFavorito := "futebol";{

	case EsporteFavorito == "vôlei":
		fmt.Printf("meu esporte favorito é vôlei\n")

	case EsporteFavorito == "golfe":
		fmt.Printf("meu esporte favorito é golfe\n")

	case EsporteFavorito == "futebol":
		fmt.Printf("meu esporte favorito é futebol\n")
	}
}