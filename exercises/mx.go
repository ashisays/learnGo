package main
 
import (
	"fmt"
	"net"
)
 
func main() {
	mxrecords, _ := net.LookupMX("google.com")
	for _, mx := range mxrecords {
		fmt.Println(mx.Host, mx.Pref)
	}
}