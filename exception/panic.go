// Simple Go program which illustrates 
// the concept of panic 
package main 

import "fmt"

// Main function 
func main() { 

	// Creating an array of string type 
	// Using var keyword 
	var myarr [3]string 

	// Elements are assigned 
	// using an index 
	myarr[0] = "panic"
	myarr[1] = "go panics"
	myarr[2] = "Gopher"

	// Accessing the elements 
	// of the array 
	// Using index value 
	fmt.Println("Elements of Array:") 
	fmt.Println("Element 1: ", myarr[0]) 

	// Program panics because 
	// the size of the array is 
	// 3 and we try to access 
	// index 5 which is not 
	// available in the current array, 
	// So, it gives an runtime error 
	fmt.Println("Element 2: ", myarr[5]) 

} 
