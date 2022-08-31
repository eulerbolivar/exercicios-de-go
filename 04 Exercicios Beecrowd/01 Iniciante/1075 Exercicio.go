package main
import "fmt"

var n int

func main() {
    fmt.Scanf("%d", &n)
    
    for i := 0; i < 10000; i++{
        if (i % n) == 2{
            fmt.Printf("%d\n", i)
        }
    }
}   