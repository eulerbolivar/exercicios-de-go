package main
import "fmt"

var x, y int

func main() {
    
    fmt.Scanf("%d %d", &x, &y)
    
    var i int
    for i < y{
        for j := 0; j < x; j++{
            i++
            if (i % x) != 0{
                fmt.Printf("%d ", i)
            } else{
                fmt.Printf("%d", i)
            }
        }
        fmt.Printf("\n")
    }
    
}