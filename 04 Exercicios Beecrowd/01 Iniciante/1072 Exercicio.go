package main

import (
    "fmt"
)

var n, i, in, out int

func main() {
    fmt.Scanf("%d", &n)

    var x = make([]int, n)

    for i = 0; i < n; i++{
        fmt.Scanf("%d", &x[i])
    }
    
    out = n
    for i = 0; i < n; i++{
        if x[i] >= 10 && x[i] <= 20{
            in += 1
            out -= 1
        }
    }
    
    fmt.Printf("%d in\n", in)
    fmt.Printf("%d out\n", out)
}