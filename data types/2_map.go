/*
A map is a builtin type in Go that is used to store key-value pairs. 

-- A map can be created by passing the type of key and value to the make function. The following is the syntax to create a new map.

make(map[type of key]type of value)  

-- It is also possible to initialize a map during the declaration itself.
  var employeeSalary = map[string]int {
        "steve": 12000,
        "jamie": 15000,
    }

-- Deault value of map is nil

-- Checking if a key exists ; value, ok := map[key]  

-- Iterate over all elements in a map;
   for key, value := range map

-- Deleting items from a map
   delete(map, key)

-- Maps are reference types
   Similar to slices, maps are reference types. When a map is assigned to a new variable, they both point to the same internal data structure.

-- Maps equality; can't be compared using the == operator. The == can be only used to         check if a map is nil.

--Performance and implementation
  - Maps are backed by hash tables.
  - Add, get and delete operations run in constant expected time. The time complexity for     the add operation is amortized.
  - The comparison operators == and != must be defined for the key type. 
*/

package main

import "fmt"

func main() {
var m map[string]float64
m = make(map[string]float64)

m["pi"] = 3.14             // Add a new key-value pair
m["pi"] = 3.1416           // Update value
fmt.Println(m)             // Print map: "map[pi:3.1416]"
// fmt.Println("length of map",len(m))  

// // fmt.Println("\nIteration over map:")
for key, value := range m { // Order not specified 
     fmt.Println(key, value)
}

v := m["pi"]               // Get value: v == 3.1416
fmt.Println("\nfetching pi",v)
v1 := m["pie"]               // Not found: v == 0 (zero value)
fmt.Println("fetching pie",v1)

if _ , found := m["pi"]; found {
     fmt.Println("found",found)
}                           // Prints "3.1416"

delete(m, "pi")             // Delete a key-value pair
fmt.Println("map \n",m)              // Print map: "map[]"

}
