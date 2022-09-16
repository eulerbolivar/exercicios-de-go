package main
import "fmt"

var qtd, num int

func main() {
    fmt.Scanf("%d", &qtd)
    
    for i := 0; i < (qtd * 2); i++{
        if i % 2 == 0{
            num++
        }
        
        if i % 2 == 0{
            fmt.Printf("%d %d %d\n", num, (num * num), (num * num * num))
            
        } else if i % 2 != 0{
            fmt.Printf("%d %d %d\n", num, ((num * num) + 1), ((num * num * num) + 1))
        }
    }
}