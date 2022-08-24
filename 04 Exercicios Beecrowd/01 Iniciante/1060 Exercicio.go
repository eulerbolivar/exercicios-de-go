package main
import "fmt"

var array [6]float64
var positivos int

func main() {
    for i := 0; i < 6; i++{
        fmt.Scanf("%f", &array[i])   
    }
    
    for i := 0; i < 6; i++{
        if (array[i] > 0){
            positivos += 1
        }
    }
        
    fmt.Printf("%d valores positivos\n", positivos)
}