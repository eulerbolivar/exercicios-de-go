package main
 
import (
    "fmt"
)
 
func main() {

    var n, x, y, somado, k int

    fmt.Scanf("%d", &n)
    for i := 0; i < n; i++{
        fmt.Scanf("%d %d", &x, &y)

        for j := x; k < y; j++{
            if (j % 2) != 0{
                somado += j
                k++
            }
        }
        fmt.Printf("%d\n", somado)
        somado = 0
        k = 0
    }
    
}
