package main
import "fmt"

var idade, contador, somados int

func main() {
    
    for i := 0; i != 1;{
        fmt.Scanf("%d", &idade)
        
        if idade > 0{
           somados += idade
        contador += 1 
        } else{
            i = 1
        }
    }
    
    media := float64(somados) / float64(contador)
    fmt.Printf("%.2f\n", media)
    
}