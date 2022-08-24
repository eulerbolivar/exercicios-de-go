package main
import "fmt"

var mes int

func main() {
    fmt.Scanf("%d", &mes)
    
    switch{
    case mes == 1:
        fmt.Printf("January\n")
    case mes == 2:
        fmt.Printf("February\n")
    case mes == 3:
        fmt.Printf("March\n")
    case mes == 4:
        fmt.Printf("April\n")
    case mes == 5:
        fmt.Printf("May\n")
    case mes == 6:
        fmt.Printf("June\n")
    case mes == 7:
        fmt.Printf("July\n")
    case mes == 8:
        fmt.Printf("August\n")
    case mes == 9:
        fmt.Printf("September\n")
    case mes == 10:
        fmt.Printf("October\n")
    case mes == 11:
        fmt.Printf("November\n")
    case mes == 12:
        fmt.Printf("December\n")
    }
}