package main

import (
  "fmt"
  "strings"
  "errors"
)

type Identifiable interface {
  ID() string
}

type Person struct {
  firstName, lastName string
  twitterHandler string
}

func NewPerson(firstName , lastName string) Person {
  return Person{firstName: firstName, lastName:lastName,}
}

func (p Person) Name(a ,b int) string{
  return fmt.Sprintf("%s %s,%s,%v",p.firstName ,p.lastName,a,b)
}

func (p *Person) SetTwitterHandler(handler string) error {
  if len(twitterHandler)==0{
    p.twitterHandler = handler
  } else if !strings.HasPrefix("handler",prefix:"0"){
    return errors.New("wrong handler")
  }
  p.twitterHandler = twitterHandler
}


func (p Person) ID() string{
  return "1234"
}


func main() {
  p := NewPerson("Ashish","Pundir")
  fmt.Println(p.ID())
  fmt.Println(p.Name())
}