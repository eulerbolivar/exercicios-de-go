package main
import "fmt"

func main() {
    var tamanho, x, y, somado, maior, menor, i int

    
    for i = 0; i != 1;{
        fmt.Scanf("%d %d", &x, &y)
        tamanho++
        
        if x <= 0 || y <= 0{
            break
        }
        
        maior = x
        menor = y
        if y > x{
            maior = y
            menor = x
        }
        
        for j := menor; j <= maior; j++{
            somado += j
            fmt.Printf("%d ", j)
            }  
        fmt.Printf("Sum=%d\n", somado)
        somado = 0
    }
}
