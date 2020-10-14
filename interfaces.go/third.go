package main
import (
 "fmt"
 "reflect"
)
type User struct {
 UserEmail string
 UserPass  string
}
func login(user interface{}) {
 
 ////--- Extract Value without specifying Type
 val := reflect.ValueOf(user).Elem()
 n := val.FieldByName("UserEmail").Interface().(string)
 fmt.Printf("%+v\n", n)
 
 fmt.Println(user.(*User).UserEmail)
}

func main() {
 login(&User{UserEmail: "lucian@knesk.com", UserPass: "lucian123"})
}