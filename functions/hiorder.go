package main
import "fmt"

func apply(x, y int, add func(int, int) int, sub func(int, int) int) (int, int) {
    r1 := add(x, y)
    r2 := sub(x, y)
    return r1, r2
}


func main() {
    x := 3
    y := 4
    add, sub := getAddSub()
    r1, r2 := apply(x, y, add, sub)
    fmt.Printf("%d + %d = %d\n", x, y, r1)
    fmt.Printf("%d - %d = %d\n", x, y, r2)
}

func getAddSub() (func(int, int) int, func(int, int) int) {
    add := func(x, y int) int {
        return x + y
    }
    sub := func(x, y int) int {
        return x - y
    }
    return add, sub
}