package main

import "fmt"

var n1, n2, n3, n4, media, exame, ne float64

func main() {
    fmt.Scanf("%f %f %f %f", &n1, &n2, &n3, &n4)
    
    n1 = (n1 * 2.0)
    n2 = (n2 * 3.0)
    n3 = (n3 * 4.0)
    n4 = (n4 * 1.0)
    
    media = (n1 + n2 + n3 + n4) / 10
    
    fmt.Printf("Media: %.1f\n", media)
    
    if media >= 7 {
        fmt.Printf("Aluno aprovado.\n")
    } else if media >= 5 && media <= 6.9{
        fmt.Printf("Aluno em exame.\n")
        fmt.Scanf("%f", &exame)
        fmt.Printf("Nota do exame: %.1f\n", exame)
        ne = (media + exame) / 2
        
        if ne >= 5 {
            fmt.Printf("Aluno aprovado.\n")
            fmt.Printf("Media final: %.1f\n", ne)
        } else{
            fmt.Printf("Aluno reprovado.\n")
        }
        
    } else if media <= 4.9 {
        fmt.Printf("Aluno reprovado.\n")
    }
     
}