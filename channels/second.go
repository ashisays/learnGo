/*
a is channel
<- a
a <- 10
data := <- a // read from channel a  
a <- data // write to channel a  
*/
package main

import (  
    "fmt"
)

func Greet(done chan bool) {  
    fmt.Println("Hello world goroutine")
    done <- true
}
func main() {  
    done := make(chan bool)
    go Greet(done)
    fmt.Println(<-done)
    fmt.Println("main function")
}