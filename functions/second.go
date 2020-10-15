package main
import "fmt"
func Mul(nums ...int) int {
   res := 0
   for _, n := range nums {
       res += n
   }
   return res
}

func main(){
   fmt.Println(Mul())        // 0
   fmt.Println(Mul(1, 2, 3)) // 6
}