package main

import(
	"fmt"
)

func main(){

	tam := 10
	var aux int

	slice := make([]int, tam)

	for i := 0; i < tam; i++{
		fmt.Scanf("%d", &slice[i])
	}

	for i := 0; i < tam; i++{
		for j := 0; j < tam - 1; j++{
			if slice[j] > slice[j+1]{
				aux = slice[j]
				slice[j] = slice[j+1]
				slice[j+1] = aux
			}
		}
	}

	fmt.Printf("nosso slice reorganizado ficou: %d\n", slice)

}