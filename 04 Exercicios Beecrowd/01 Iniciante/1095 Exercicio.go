package main
import "fmt"

func main() {
    var j int = 60
    
    for i := 1; j >= 0; i += 3{
        fmt.Printf("I=%d J=%d\n", i, j)
        j -= 5
    }
}