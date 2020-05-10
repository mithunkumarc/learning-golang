package main

import (
	"fmt"
	"reflect"
)

func main() {
	//functional prog: returning functions, object functions treated equally
	f := bar() // return anonymous function
	fmt.Println(f()) // 42
	fmt.Println(reflect.TypeOf(f)) // func() int
}

// bar function has return type function which returns int
func bar() func() int {
	return func() int {
		return 42
	}
}
