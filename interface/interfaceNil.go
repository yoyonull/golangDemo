package main

import (
	"fmt"
	"reflect"
)

func main() {
	t := []string{"stanley", "david", "oscar"}
	s := make([]interface{}, len(t))
	for i, v := range t {
		s[i] = v
	}
	fmt.Println(s)
	fmt.Println("type:", reflect.TypeOf(s))
	for _, v := range s {
		fmt.Println(v)
	}
	//var value interface{} // Value provided by caller.
	//str, ok := value.(string)
	//if ok {
	//	fmt.Printf("string value is: %q\n", str)
	//} else {
	//	fmt.Printf("value is not a string\n")
	//}
}
