package main
 
import (
    "fmt"
)

var a float64
var b float64

func main() {
    fmt.Scanf("%f", &a)
    fmt.Scanf("%f", &b)

    a = (a * 3.5)
    b = (b * 7.5)

    media := (a/11) + (b/11)
    fmt.Printf("MEDIA = %.5f\n", media)
}
