package main
import "fmt"

var n int

func main() {
    fmt.Scanf("%d", &n)
    
    v1 := make([]float64, n)
    v2 := make([]float64, n)
    v3 := make([]float64, n)
    media := make([]float64, n)
    
    for i := 0; i < n; i++{
        fmt.Scanf("%f %f %f", &v1[i], &v2[i], &v3[i])
    }
        
    for i := 0; i < n; i++{
        media[i] = ((v1[i] * 2) + (v2[i] * 3) + (v3[i] * 5))/10
        fmt.Printf("%.1f\n", media[i])
    }
}   
