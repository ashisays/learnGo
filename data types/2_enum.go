/* enumeration
  ENUMs are a powerful feature of many languages. They let you define strict values of data you expect.
  
  With Go, it’s different. Most often you create a custom type (enum_name) and define  constants of that type. In the case of integer custom type, iota keyword can be used to simplify the definition.

  1. Using Constant Declaration with iota
    Within a constant declaration, the predeclared identifier iota represents successive untyped integer constants
    const (  // iota is reset to 0
        c0 = iota  // c0 == 0
        c1 = iota  // c1 == 1
        c2 = iota  // c2 == 2
    )

  2. const with Expression.
    const (
        a = 1 << iota  // a == 1 (iota has been reset)
        b = 1 << iota  // b == 2
        c = 1 << iota  // c == 4
    )

    const (
        u         = iota * 42  // u == 0     (untyped integer constant)
        v float64 = iota * 42  // v == 42.0  (float64 constant)
        w         = iota * 42  // w == 84    (untyped integer constant)
    )

  3. Using Struct.
    - const have a lot boiler plate code and runtime issues.
    - Enums limit what values a type can hold. So really, enums are just an extension on  what a type can do
    
    type Alias = int

    type list struct { 
    West Alias
    East Alias
    North Alias
    South Alias
  }

  // Enum for public use
  var Enum = &list{ 
    West: 0,
    East: 1,
    North: 2,
    South: 3,
  }    

*/

package main

import "fmt"

// A Month specifies a month of the year (January = 1, ...).
type Month int

const (
 January Month = 1 + iota
 February
 March
 April
 May
 June
 July
 August
 September
 October
 November
 December
)

//slices of  string for month
var months = [...]string{
 "January",
 "February",
 "March",
 "April",
 "May",
 "June",
 "July",
 "August",
 "September",
 "October",
 "November",
 "December",
}

// String returns the English name of the month mapping const to string
//("January", "February", ...).
func (m Month) String() string { return months[m-1] }


// Struct as the enumeration
// Alias hide the real type of the enum 
// and users can use it to define the var for accepting enum 
type Alias = int

type list struct { 
    West Alias
    East Alias
    North Alias
    South Alias
}

// Enum for public use
var Enum = &list{ 
    West: 0,
    East: 1,
    North: 2,
    South: 3,
}

func main() {
 month := December
 if month == December {
  fmt.Println("Found a December")
 }
 // %!v(PANIC=runtime error: index out of range)
 month = month + Month(2)
//  fmt.Println(month)

 month = January + Month(2)
 fmt.Println(month)

 month++
 fmt.Println(month)

 day := 34
 month = Month(day % 31)
 fmt.Println(month)

 //check month value
 val := int(month) + 4
 fmt.Println(val)

 month = Month(val) + 1
 fmt.Println(month)

  fmt.Println("\n Struct as Enum Example !!")
  wayToGo := Enum.West
    fmt.Println("Where to go? ", wayToGo)
}
