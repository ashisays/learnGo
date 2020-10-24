package main
import "fmt"

func f1(c chan int, x int) {
fmt.Println(x)
c <- x
}
func f2(c chan<- int, x int) {
fmt.Println(x)
c <- x
}
func main () {
  c := make(chan int)
  go f1(c,1)
  fmt.Println("REad 1: ",<-c)
  go f2(c,2)
  fmt.Println("REad 2: ",<-c)
}