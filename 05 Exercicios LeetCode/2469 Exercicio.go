package main

import(
	"fmt"
)

// FUNÇÃO SOLICITADA PELO EXERCÍCIO
func convertTemperature(celsius float64) []float64{
	
	kelvin := celsius + 273.15
	fahrenheit := (celsius * 1.8) + 32.0

	ans := make([]float64, 2)
	ans[0] = kelvin
	ans[1] = fahrenheit

	return ans
}

// ENTRADA PADRÃO PARA O USO DA FUNÇÃO
func main(){

	resp := convertTemperature(122.11)

	fmt.Printf("%.2f\n", resp)
}