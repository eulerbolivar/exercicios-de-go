package main
import "fmt"

var n, i int

func main() {
    fmt.Scanf("%d", &n)
    
    slice := make([]int, n)
    
    for i = 0; i < n; i++{
        fmt.Scanf("%d", &slice[i])
    }
    
    for i = 0; i < n; i++{
        if (slice[i] == 0){
            fmt.Printf("NULL\n")
        }
            
        if ((slice[i] % 2) == 0 && slice[i] != 0){
            fmt.Printf("EVEN ")
            
            if (slice[i] > 0){
                fmt.Printf("POSITIVE\n")
            }
            if (slice[i] < 0){
                fmt.Printf("NEGATIVE\n")
            }
        }
            
        if ((slice[i] % 2) != 0){
            fmt.Printf("ODD ")
            
            if (slice[i] > 0){
                fmt.Printf("POSITIVE\n")
            }
            if (slice[i] < 0){
                fmt.Printf("NEGATIVE\n")
            }
        }
    }
}