//unidirectional channel
package main

import "fmt"

func sendData(sendch chan<- int) (b <-chan int) { 
    a := <-sendch
    fmt.Println(a) 
    b <- 10
    return 
}

func main() {  
    sendch := make(chan int)
    c := sendData(sendch)
    fmt.Println(<-sendch)
    fmt.Println(sendch)
}