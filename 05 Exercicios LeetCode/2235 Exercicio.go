package main

import (
	"fmt"
)

func sum(num1 int, num2 int) int {
    ans := num1 + num2

	return ans
}


func main(){
	var x, y, resp int

	fmt.Scanf("%d", &x)
	fmt.Scanf("%d", &y)

	resp = sum(x, y)

	fmt.Printf("%d\n", resp)
}