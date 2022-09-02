package main
import "fmt"

func main() {
    var j int = 7
    var c int
    
    for i := 1; i <= 9;{
        
        fmt.Printf("I=%d J=%d\n", i, j)
        j -= 1
        c += 1
        
        if c == 3 {
            j += 5
            i += 2
            c = 0
        }
    }
}