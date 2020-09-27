package main

import "fmt"

func main() {
	fmt.Println("This module print febonacci number")
	fmt.Println("Need to find the fibonacci for no. =", 10)
  fibonacci(10)
}

func fibonacci(a int) int {
	p := 0
	q := 1
	for x :=1; x <= a; x++ {
		temp := p
    p = p + q
    q = temp
    fmt.Printf("%d ,",p)
	}
  fmt.Println("")
	return 0
}
