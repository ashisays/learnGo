package main
 
import (
	"fmt"
	"net"
)
 
func main() {
	mxrecords, _ := net.LookupMX("botcliq.tech")
	for _, mx := range mxrecords {
		fmt.Println(mx.Host, mx.Pref)
	}
}