package main 
  
import (
  "fmt"
  "reflect"
)

// Finalsalary of function type 
type Finalsalary func(int, int) int
  
// Creating structure 
type Author struct { 
    name      int 
    language  int 
    Marticles int
    Pay       int
  
    // Function as a field 
    //salary Finalsalary 
} 
  
// Main method 
func main() { 
  
    // Initializing the fields 
    // of the structure 
    result := Author{ 
        name:      1, 
        language:  1, 
        Marticles: 120, 
        Pay:       500,  
    } 
  
    // Display values 
    fmt.Println("Author's Name: ", result.name) 
    fmt.Println("Language: ", result.language) 
    fmt.Println("Total number of articles published in May: ", result.Marticles) 
    fmt.Println("Per article pay: ", result.Pay) 
    fmt.Printf("Size of var (reflect.TypeOf.Size): %d\n", reflect.TypeOf(result).Size())
    //fmt.Println("Total salary: ", result.salary(result.Marticles, result.Pay)) 
}   