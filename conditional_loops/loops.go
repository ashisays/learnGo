package main 
  
import "fmt"
  
// Main function 
func main() { 
   fmt.Println("1. For loop Default Format")   
    // for loop  
    // This loop starts when i = 0  
    // executes till i<4 condition is true 
    // post statement is i++ 
    for i := 0; i < 4; i++{ 
      fmt.Println("iteration",i)   
    } 
    
    fmt.Println("\n For as infinite loop, terminate at 10")
    j := 0
    for { 
      fmt.Println("infinite",j)
      if j == 10{
        fmt.Println("Stopping infinite loop using break !!")
        break
      }   
      j++
    }

    fmt.Println("\n For as While loop, terminate at 10")
    j = 0
    for j < 10{ 
      fmt.Println("while j < 10",j)
      j++   
    }
    
    fmt.Println("\n For in range")
    for i, j:= range "Hello Go Learner !! Enjoy" { 
       fmt.Printf("The index number of %U is %d and %c\n", j, i,j)  
    } 
    
} 