/*
  1. Create an Interface with Method Calculate Area.

  2. Create Square, rectangle and Triangle struct, will have values for fields and implement Calculate Area Function.

  3. Main Function
    - intialize the value of all structure.
    - CalculateArea and print value.

*/

package main

import (
	"fmt"
)

type shaper interface {
  CalculateArea() float64
}

type square struct {
	length float64
}

func (s square) CalculateArea()float64 {
  return s.length*s.length
}


type rectangle struct {
	width  float64
	height float64
}

func (r rectangle) CalculateArea()float64 {
  return r.width*r.height
}

type triangle struct {
  base float64
  height float64
}

func (t triangle) CalculateArea() float64 {
  return (t.base*t.height)/2
}

func main (){
  var si []shaper
  si = append(si,square{length:10})
  si = append(si,rectangle{width:10,height:20})
  si = append(si,triangle{base:20,
        height:100})
  fmt.Println("Iterating over interface:")
  var max float64
  var i shaper
  for _,v := range si {
    fmt.Printf("Type: %T, Value %#v\n",v,v)
    area := v.CalculateArea()
    if area > max{
      max = area
      i = v
    }
  }
  fmt.Printf("Max Area in the Boxes is :\n\t Box: %#v,\n\t Area : %v\n",i,max)
}