package main
import "fmt"

func factorial(x int64) int64{
  var fact int64 =1
  var i int64 = 1
  for ;i <= x; i++ {
     fact = fact*i
    fmt.Println("factorial", fact , i)
  }
  fmt.Println("factorial", fact)
  return fact
}

func main() {
  val := factorial(5)
  fmt.Printf("factorial of 5 is %d \n",val)
  val = factorial(20)
  fmt.Printf("factorial of 20 is %d \n",val)
  fmt.Println()
}
