package main
 
import (
    "fmt"
)
var A int
var B int 

func main() {
    fmt.Scanf("%v", &A)
    fmt.Scanf("%v", &B)

    X := A + B

    fmt.Printf("X = %v\n", X)
}
