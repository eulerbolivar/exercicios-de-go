package main
import "fmt"

func main() {
    var j int = 7
    
    for i := 1; i <= 9;{
        fmt.Printf("I=%d J=%d\n", i, j)
        j -= 1
        if j < 5{
            j = 7
            i += 2
        }
    }
}