package main

import (
    "fmt"
)

var n, i int

func main() {
    fmt.Scanf("%d", &n)

    pares := 2
    for i = 0; i < n; i += 2{
        fmt.Printf("%d^2 = %d\n", pares, (pares * pares))
        pares += 2
    }
}