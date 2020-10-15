package main
import "fmt"
func Mul(nums []int) int {
   res := 0
   for _, n := range nums {
       res += n
   }
   return res
}

func main(){
   fmt.Println(Mul([]int{}))        // 0
   fmt.Println(Mul([]int{1, 2, 3})) // 6
}