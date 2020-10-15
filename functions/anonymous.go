package main

    import "fmt"

    func main() {

    sum := func(a, b, c int) int {
        return a + b + c
    }
    fmt.Println(sum(5,3,7))
    }