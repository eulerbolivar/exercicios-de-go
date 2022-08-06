package main
 
import (
    "fmt"
)

var a float64
var b float64
var c float64


func main() {
    fmt.Scanf("%f", &a)
    fmt.Scanf("%f", &b)
    fmt.Scanf("%f", &c)

    a = (a * 2)
    b = (b * 3)
    c = (c * 5)

    media := (a/10) + (b/10) + (c/10)

    fmt.Printf("MEDIA = %.1f\n", media)
}
