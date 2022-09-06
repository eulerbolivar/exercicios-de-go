package main
 
import (
    "fmt"
)
 
var x, y int

func main() {

    for i := 0; i != 1;{
        fmt.Scanf("%d %d", &x, &y)

        if x == y{
            i = 1
        }

        if x < y{
            fmt.Printf("Crescente\n")
        } else if y < x{
            fmt.Printf("Decrescente\n")
        }
    }

    
}
