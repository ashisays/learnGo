package main

import "fmt"

func goroutineFunc() {
    fmt.Println("Inside my goroutine")
}

func main() {
    fmt.Println("Hello World")
    go goroutineFunc()
    fmt.Println("Finished Execution")
}