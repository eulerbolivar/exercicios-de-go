package main

import (
    "fmt"
    "math"
    )
    
var a, b, c, delta, x1, x2 float64

func main() {
    fmt.Scanf("%f %f %f", &a, &b, &c)
    
    delta = (b * b) - (4 * a * c)
    
    if delta < 0 || a == 0{
        fmt.Printf("Impossivel calcular\n")
    }else{
        x1 = (-b + math.Sqrt(delta)) / (2 * a)
        x2 = (-b - math.Sqrt(delta)) / (2 * a)
        
        fmt.Printf("R1 = %.5f\n", x1)
        fmt.Printf("R2 = %.5f\n", x2)
    }
}