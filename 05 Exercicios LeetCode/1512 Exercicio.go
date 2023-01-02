package main

import(
	"fmt"
)

// FUNÇÃO SOLICITADA PELO EXERCÍCIO
func numIdenticalPairs(nums []int) int {
	
	var ans int

	for i := 0; i < len(nums); i++{
		for j := 1 + i; j < len(nums); j++{
			if nums[i] == nums[j]{
				ans += 1
				//fmt.Printf("%d e %d são iguais!\n\n", nums[i], nums[j])
			}
		}
	}

	return ans
}

// ENTRADA PADRÃO PARA O USO DA FUNÇÃO
func main(){
	intArray := []int{1, 1, 1, 1}
	resp := numIdenticalPairs(intArray)
	fmt.Printf("%d\n", resp)

}