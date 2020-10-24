package main

import "fmt"

func main() {  
    var a chan int
    fmt.Println(a)
    a = make(chan int)
    //a := make(chan int)
    fmt.Println(a)
    fmt.Printf("Type of a is %T", a)
}