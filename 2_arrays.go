/*
  ARrays:
  - array in golang is mutable.
  - Create array using two syntax.
    Var array_name[length]Type
    or
    var array_name[length]Typle{item1, item2, item3, ...itemN}
  - iterage using range in for loop.

*/
package main 
  
import "fmt"
  
func main() { 
  
// Creating an array of string type  
// Using var keyword 
var myarr[3]string
  
// Elements are assigned using index 
myarr[0] = "Hello"
myarr[1] = "Go"
myarr[2] = "World !!"
  
// Accessing the elements of the array  
// Using index value 
fmt.Println("Elements of Array:") 
fmt.Println("Element 1: ", myarr[0]) 
fmt.Println("Element 2: ", myarr[1]) 
fmt.Println("Element 3: ", myarr[2])

for _, value := range myarr{
  fmt.Printf("%v ",value)
}
fmt.Println("")

// Creating and initializing  
// 2-dimensional array  
// Using shorthand declaration 
// Here the (,) Comma is necessary 
arr:= [3][3]string{{"C#", "C", "Python"},  
                   {"Java", "Scala", "Perl"}, 
                    {"C++", "Go", "HTML"},} 

// Accessing the values of the  
// array Using for loop 
fmt.Println("\n Elements of MultiDimensional Array") 
for x:= 0; x < 3; x++{ 
  for y:= 0; y < 3; y++{ 
    fmt.Printf("%s ,",arr[x][y]) 
  } 
  fmt.Printf("\n")
} 

} 