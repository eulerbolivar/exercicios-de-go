package main

import (
	"fmt"
)

func getConcatenation(nums []int) []int {
    n := len(nums)
    ans := make([]int, n*2)
    var c int = 0
    
    for i := 0; i < 2*n; i++{
        ans[i] = nums[c]
        c++
        if c == n{
            c = 0
        }
    }
    
    return ans
}

func main(){
	numeros := make([]int, 3)

	for i := 0; i < 3; i++{
		fmt.Scanf("%d", &numeros[i]) 
	}

	fmt.Printf("%d\n", getConcatenation(numeros[:]))
}