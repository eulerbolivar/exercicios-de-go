package main
import "fmt"

func main() {
    var quantia, x, y, somado, maior, menor int
    
    fmt.Scanf("%d", &quantia)
    
    for i := 0; i < quantia; i++{
        fmt.Scanf("%d %d", &x, &y)
        
        
        maior = x
        menor = y
        if y > x{
            maior = y
            menor = x
        }
        
        
        for j := menor + 1; j < maior; j++{
            if (j % 2) != 0{
                somado += j
            }
        }   
        
        fmt.Printf("%d\n", somado)
        somado = 0
    }
}