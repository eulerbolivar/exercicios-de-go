package main

import(
	"fmt"
)

func shuffle(nums []int, n int) []int {
	var k, j int
    ans := make([]int, len(nums))

	// LOOP DO SHUFLE DISTRIBUINDO 
	for i := 0; i != len(nums); i++{
		if (i % 2) == 0{
			ans[i] = nums[j]
			j++
		}else {
			ans[i] = nums[(len(nums)/2)+k]
			k++
		}
	}

	return ans
}

func main(){
	var tam int

	// DEFINE O TAMANHO DO ARRAY
	fmt.Scanf("%d", &tam)

	// DECLARA O SLICE E COLOCA OS VALORES NELE
	slice := make([]int, (tam * 2))
	for i := 0; i < (tam * 2); i++{
		fmt.Scanf("%d", &slice[i])
	}

	resp := shuffle(slice, tam)

	fmt.Printf("%d\n", resp)

}