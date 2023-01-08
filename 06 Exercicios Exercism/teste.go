package main

import (
	"fmt"
)

func main(){
	array := []int{1,2,3,4,5}

	fmt.Printf("%d\n", array[len(array) - 1:])
	fmt.Printf("%d\n", array[:len(array) - 1])
}