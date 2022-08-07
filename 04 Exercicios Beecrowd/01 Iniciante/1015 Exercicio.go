package main
 
import (
    "fmt"
    "math"
)

var x1 float64
var y1 float64

var x2 float64
var y2 float64

func main() {
    fmt.Scanf("%f %f", &x1, &y1)
    fmt.Scanf("%f %f", &x2, &y2)

    distancia := math.Sqrt((x2 - x1) * (x2 - x1) + (y2 - y1) * (y2 - y1))

    fmt.Printf("%.4f\n", distancia)
}
