package main

import (
   "fmt"
)

func main() {
   id, err := ReturnId()

   if err != nil {
      fmt.Printf("ERROR: %s", err)
      return
   }

   fmt.Printf("Id: %d\n", id)
}

func ReturnId() (id int,err  error) {
   id = 10
   if id == 10 {
      err = fmt.Errorf("Invalid Id\n")
      return
   }
   return
}