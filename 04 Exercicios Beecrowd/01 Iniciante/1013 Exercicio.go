package main
 
import (
    "fmt"
)

var x int
var y int 
var z int 

func main() {
    fmt.Scanf("%v %v %v", &x, &y, &z)

    maior := x
    
    if (y > maior){
        maior = y
    }

    if (z > maior){
        maior = z
    }

    fmt.Printf("%v eh o maior\n", maior)
}
