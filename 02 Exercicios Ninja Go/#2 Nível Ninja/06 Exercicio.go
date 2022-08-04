package main

import (
	"fmt"
)

const (
	_ = 1917 + iota
	a
	b
	c
	d
)

func main() {
	fmt.Println(a, b, c, d)
}