package main

import(
	"fmt"
)

func teste()int{
	ans := 5

	return ans
}

func main(){

	resp := teste()
	fmt.Printf("%d\n", resp)

}