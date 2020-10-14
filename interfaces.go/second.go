package main

import (
  "fmt"
  "reflect"
)

// interface definition
type Publisher interface {
     Publish()  error
}

type blogPost struct {
  Author  string
  title   string
  postId  int  
}

// method with a value receiver
func (b blogPost) Publish() error {
   fmt. Printf("The title on %s has been published by %s, with postId %d\n" , b.title, b.Author, b.postId)
   return nil
}

 func test(){
  b := blogPost{"Alex","understanding structs and interface types",12345}
  fmt.Println(b.Publish())
   d := &b   // pointer receiver for the struct type
   b.Author = "Chinedu"
   fmt.Println(d.Publish())

}

func main() {
        var p Publisher
        fmt.Println(p)
        p = blogPost{"Alex","understanding structs and interface types",12345}
        fmt.Println(p.Publish())
        fmt.Println("\n********************************\n")
        fmt.Println("Blog is %v",p)
        fmt.Println("\n********************************\n")
        c := reflect.ValueOf(p).Elem()
        fmt.Println("Author is : ", c.Author)
        fmt.Println("\n********************************\n")
        test()  // call the test function 

}