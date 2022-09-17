package main
 
import (
    "fmt"
)
 
func main() {
    var num int
    fmt.Scanf("%d", &num)

    for i := 1; i < num + 1; i++{
        if num % i == 0{
            fmt.Printf("%d\n", i)
        }
    }
}
