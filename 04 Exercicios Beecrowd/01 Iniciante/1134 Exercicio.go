package main
 
import (
    "fmt"
)

var tipoCombust int
var alcool, gasolina, diesel int

func main() {
    
    for i := 0; i != 2;{
        fmt.Scanf("%d", &tipoCombust)

        switch {
        case tipoCombust == 1:
            alcool += 1
        case tipoCombust == 2:
            gasolina += 1
        case tipoCombust == 3:
            diesel += 1
        case tipoCombust == 4:
            i = 2   
        }
    }
    fmt.Printf("MUITO OBRIGADO\n")
    fmt.Printf("Alcool: %d\nGasolina: %d\nDiesel: %d\n", alcool, gasolina, diesel)
}
    
