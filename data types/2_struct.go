/*
A struct or structure can be compared with the class in the Object-Oriented Programming paradigm

-- A structure has different fields of the same or different data types.

-- We use struct keyword to create a new structure type as shown in the example below.
   type StructName struct {
    field1 fieldType1
    field2 fieldType2
  }

-- Getting and setting struct fields
   When a struct variable is created, we can access its fields using . (dot) operator
   e.g. StructName.field1

-- An anonymous struct is a struct with no explicitly defined derived struct type.
   ross := struct{
          firstName string
          lastName string
          fullTime bool
    } {
    firstName: "ross",
    lastName:  "Bing",
    fullTime:  true,
  }

  -- The syntax to create a pointer to a struct is as follows.
    s := &StructType{...}

  -- Anonymous fields; You can define a struct type without declaring any field names. You have to just define the field data types and Go will use the data type declarations (keywords) as the field names.
    type Employee struct {
	  firstName, lastName string
	  salary              int
	  bool                         // anonymous field
    }
  
  -- A struct field can be of any data type. Hence, it is perfectly legal to have a struct field that holds another struct. 

  -- Like a struct, an interface can also be nested in a struct. In Layman’s terms, it means that a field can have a data type of an interface.

  -- struct fields can also be functions.

  -- Two structs are comparable if they belong to the same type and have the same field values.

  -- Struct gives one more ability to add meta-data to its fields. 
     type Employee struct {
    	firstName string `json:"firstName"`
    	lastName  string `json:"lastName"`
    	salary    int    `json: "salary"`
    	fullTime  int    `json: "fullTime"`
    } 
*/
package main

import "fmt"

type Salary struct {
	Basic, HRA, TA float64
}

type Employee struct {
	FirstName, LastName, Email string
	Age                        int
	MonthlySalary              []Salary
}

func (e Employee) EmpInfo() string {
	fmt.Println(e.FirstName, e.LastName)
	fmt.Println(e.Age)
	fmt.Println(e.Email)
	for _, info := range e.MonthlySalary {
		fmt.Println("===================")
		fmt.Println(info.Basic)
		fmt.Println(info.HRA)
		fmt.Println(info.TA)
	}
	return "----------------------"
}

func main() {

	e := Employee{
		FirstName: "Mark",
		LastName:  "Jones",
		Email:     "mark@gmail.com",
		Age:       25,
		MonthlySalary: []Salary{
			Salary{
				Basic: 15000.00,
				HRA:   5000.00,
				TA:    2000.00,
			},
			Salary{
				Basic: 16000.00,
				HRA:   5000.00,
				TA:    2100.00,
			},
			Salary{
				Basic: 17000.00,
				HRA:   5000.00,
				TA:    2200.00,
			},
		},
	}

	fmt.Println(e.EmpInfo())
}