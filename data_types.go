package main

import (
	"fmt"
  "reflect"
)

func main() {
  fmt.Printf("Float Example:")
  a := 20.45 
  b := 34.89 
  // Subtraction of two  
  // floating-point number 
  c := b-a 
    // Display the result  
  fmt.Printf("\nResult is: %f", c)       
    // Display the type of c variable 
  fmt.Printf("\nThe type of c is : %T\n", c)
	fmt.Printf("\nComplex Example:\n") 
  var c1 complex128 = complex(6, 2)
	var c2 complex64 = complex(9, 2)
	fmt.Println(c1)
	fmt.Println(c2)

	// Display the type
	fmt.Printf("The type of c1 is %T and "+
		"the type of c2 is %T \n", c1, c2)

fmt.Printf("\nRunes Example:")
  
  // Creating a rune 
    rune1 := 'B'
    rune2 := 'g'
    rune3 := '\a'
  
    // Displaying rune and its type 
    fmt.Printf("Rune 1: %c; Unicode: %U; Type: %s", rune1, 
                             rune1, reflect.TypeOf(rune1)) 
      
    fmt.Printf("\nRune 2: %c; Unicode: %U; Type: %s", rune2, 
                               rune2, reflect.TypeOf(rune2)) 
      
    fmt.Printf("\nRune 3: Unicode: %U; Type: %s\n", rune3,  
                                 reflect.TypeOf(rune3)) 
}
