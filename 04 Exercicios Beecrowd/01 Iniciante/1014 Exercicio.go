package main
 
import (
    "fmt"
)

var x float64
var y float64

func main() {
    fmt.Scanf("%f", &x)
    fmt.Scanf("%f", &y)

    media := x / y

    fmt.Printf("%.3f km/l\n", media)
}
