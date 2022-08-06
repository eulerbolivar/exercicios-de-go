package main
 
import (
    "fmt"
)

var raio float64


func main() {
    fmt.Scanf("%f", &raio)

    pi := 3.14159
    raiocubo := raio * raio * raio
    volume := (4.0/3.0) * pi * raiocubo

    fmt.Printf("VOLUME = %.3f\n", volume)

}
