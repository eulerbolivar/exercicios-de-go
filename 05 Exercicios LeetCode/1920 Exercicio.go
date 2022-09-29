package main
 
import (
    "fmt"
)

func buildArray(nums []int) []int {
    ans := make([]int, len(nums))
    for i := 0; i < len(ans); i++{
        ans[i] = nums[nums[i]]
    }
    
    return ans
}

func main(){
    numeros := make([]int, 6)

    for i := 0; i < len(numeros); i++{
        fmt.Scanf("%d", &numeros[i])
        if numeros[i] >= len(numeros){
            i--
        }
    }

    fmt.Printf("%d\n", buildArray(numeros))
}
