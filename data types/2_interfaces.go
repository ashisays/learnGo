/*

In Go language, the interface is a custom type that is used to specify a set of one or more method signatures and the interface is abstract, so you are not allowed to create an instance of the interface.

" an interface is a set of method signatures. When a type provides definition for all the methods in the interface, it is said to implement the interface."

 -- type interface_name interface{

    // Method signatures

    }

-- internal Representaion
An interface can be thought of as being represented internally by a tuple (type, value). type is the underlying concrete type of the interface and value holds the value of the concrete type.

-- The zero value of the interface is nil.
When an interface contains zero methods, such types of interface is known as the empty interface. So, all the types implement the empty interface.
Syntax:

interface{}

-- Type Assertions: In Go language, type assertion is an operation applied to the value of the interface. Or in other words, type assertion is a process to extract the values of the interface.

-- Use of Interface: 
    -You can use interface when in methods or functions you want to pass different types of argument in them just like Println () function. 
    -you can also use interface when multiple types implement same interface.

*/

// Golang program illustrates how 
// to implement an interface 
package main 
  
import "fmt"
  
// Creating an interface 
type tank interface { 
  
    // Methods 
    Tarea() float64 
    Volume() float64 
} 
  
type myvalue struct { 
    radius float64 
    height float64 
} 
  
// Implementing methods of 
// the tank interface 
func (m myvalue) Tarea() float64 { 
  
    return 2*m.radius*m.height + 
        2*3.14*m.radius*m.radius 
} 
  
func (m myvalue) Volume() float64 { 
  
    return 3.14 * m.radius * m.radius * m.height 
} 
  

func myfun(a interface{}) { 
  
    // Using type switch 
    switch a.(type) { 
  
    case int: 
        fmt.Println("Type: int, Value:", a.(int)) 
    case string: 
        fmt.Println("\nType: string, Value: ", a.(string)) 
    case float64: 
        fmt.Println("\nType: float64, Value: ", a.(float64)) 
    default: 
        fmt.Println("\nType not found") 
    } 
} 


/* empty interface
An interface that has zero methods is called an empty interface. It is represented as interface{}. Since the empty interface has zero methods, all types implement the empty interface.
*/
func describe(i interface{}) {  
    fmt.Printf("Type = %T, value = %v\n", i, i)
}


/*
Type assertion is used to extract the underlying value of the interface.

i.(T) is the syntax which is used to get the underlying value of interface i whose concrete type is 
*/
func assert(i interface{}) {  
    v, ok := i.(int)
    fmt.Println(v, ok)
}

// Main Method 
func main() { 
  
    // Accessing elements of 
    // the tank interface 
    var t tank 
    t = myvalue{10, 14} 
    fmt.Println("Area of tank :", t.Tarea()) 
    fmt.Println("Volume of tank:", t.Volume()) 

    fmt.Println("\n The value of type interface in switch !!")

myfun("GeeksforGeeks") 
    myfun(67.9) 
    myfun(true)

    fmt.Println("\nInterface Internal Representaion!!")
    fmt.Printf("Interface type %T value %v\n", t, t)

    fmt.Println("\nEmpty Interface!!")
    describe("Ashish")
    describe(77)

    fmt.Println("\n Type Assertions!!")
    assert(98)
    assert("Ashish")

} 