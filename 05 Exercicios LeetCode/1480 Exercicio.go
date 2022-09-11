package main

import (
	"fmt"
)

func runningSum(nums []int) []int {
	var sum int

    ans := make([]int, len(nums))

	for i := 0; i < len(nums); i++{
		for j := 0; j < i + 1; j++{
			sum += nums[j]
		}
		ans[i] = sum
		sum = 0
	}

	return ans
}

func main(){
	numeros := make([]int, 4)

	for i := 0; i < len(numeros); i++{
		fmt.Scanf("%d", &numeros[i]) 
	}

	fmt.Printf("%d\n", runningSum(numeros[:]))
}