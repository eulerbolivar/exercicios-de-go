package main
import "fmt"

func main() {
    var somado, x, y int
    fmt.Scanf("%d", &x)
    fmt.Scanf("%d", &y)
    
    //CONTROLE DE ORDEM CRESCENTE
    menor := x
    maior := y
    
    if y < x{
        maior = x
        menor = y
    }
    
    //LOOP PARA SOMA
    for i := menor; i <= maior; i++{
        if (i % 13) != 0{
            somado += i
        }
    }
    fmt.Printf("%d\n", somado)
}