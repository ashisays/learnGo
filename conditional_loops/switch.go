package main

import "fmt"

func main() {
	fmt.Println("Switch type Value:")
	var strvalue string = "five"

	// Switch statement without default statement
	// Multiple values in case statement
	switch strvalue {
	case "one":
		fmt.Println("C#")
	case "two", "three":
		fmt.Println("Go")
	case "four", "five", "six":
		fmt.Println("Java")
	}

	fmt.Println("Switch expression:")
	var value interface{}
	switch q := value.(type) {
	case bool:
		fmt.Println("value is of boolean type")
	case float64:
		fmt.Println("value is of float64 type")
	case int:
		fmt.Println("value is of int type")
	default:
		fmt.Printf("value is of type: %T", q)

	}
}
