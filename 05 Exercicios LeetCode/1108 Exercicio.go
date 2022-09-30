package main

import(
	"fmt"
	"strings"
)

func defangIPaddr(address string) string {
	var ans string 
	texto := strings.Split(address, "")

	// ENCONTRA E APENDA OS COLCHETES EM VOLTA DOS PONTOS
    for i := 0; i < len(texto); i++{
		if texto[i] == "."{
			texto[i-1] += "["
			texto[i] += "]"
		} 
	}

	// TRANSFORMA DE VOLTA O SLICE EM STRING PARA ENVIAR NO RETURN
	for i := 0; i < len(texto);i++{
		ans += texto[i]
	}

	return ans
}

func main(){
	var endereco string

	/*in := bufio.NewReader(os.Stdin)
	line, err := in.ReadString('\n')*/

	fmt.Scanf("%s", &endereco)

	resp := defangIPaddr(endereco)

	fmt.Printf("\n%s\n", resp)
}
