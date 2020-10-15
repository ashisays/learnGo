package main

import "fmt"

func sum(x, y, z int) (a, b int) {

    a = x + y
    b = y + z
    return
}

func main() {
    x, y := sum(10, 100,1000)
    fmt.Println(x, y)
}