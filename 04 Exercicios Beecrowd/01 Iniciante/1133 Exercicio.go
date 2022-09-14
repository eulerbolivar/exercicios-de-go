package main
import "fmt"

func main() {
    var x, y int
    
    fmt.Scanf("%d", &x)
    fmt.Scanf("%d", &y)
    
    maior := x
    menor := y
    
    if y > x{
        maior = y
        menor = x
    }
    
    for i := menor + 1; i < maior; i++{
        if (i % 5) == 2 && i != 2{
            fmt.Printf("%d\n", i)
        }
        if (i % 5) == 3 && i != 3{
            fmt.Printf("%d\n", i)
        }
    }
}