package main


import "fmt"

var valor float64

func main() {
    fmt.Scanf("%f", &valor)
    
    if valor > 0 && valor <= 25{
        fmt.Printf("Intervalo (0,25]\n")
    } else if valor > 25 && valor <= 50{
        fmt.Printf("Intervalo (25,50]\n")
    } else if valor > 50 && valor <= 75{
        fmt.Printf("Intervalo (50,75]\n")
    } else if valor > 75 && valor <= 100{
        fmt.Printf("Intervalo (75,100]\n")
    } else{
        fmt.Printf("Fora do intervalo\n")
    }
}