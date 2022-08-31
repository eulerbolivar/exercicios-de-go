package main
import "fmt"

var n, result int

func main() {
    fmt.Scanf("%d", &n)
    
    for i := 1; i <= 10; i++{
        result = i * n
        fmt.Printf("%d x %d = %d\n", i, n, result)
    }
}   