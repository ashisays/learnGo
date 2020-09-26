package main

import (
	"fmt"
	"strings"
  "os"
)

func main() {
  var originalString string 
  if len(os.Args) > 1 {
		originalString =  os.Args[1]
	} else {
    originalString = "tacocat"
	}
	
  var reverseString string = ""
	var length = len(originalString)

	for i := length - 1; i >= 0; i -- {
		reverseString = reverseString + string(originalString[i])
	}

	// Case in-sensitive comparision
	if strings.ToLower(originalString) == strings.ToLower(reverseString) {
		fmt.Println("The given string is Palindrome");
	} else {
		fmt.Println("The given string is NOT a Palindrome");
	}
}