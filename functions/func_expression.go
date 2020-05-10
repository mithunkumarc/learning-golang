package main

import (
	"fmt"
	"reflect"
)

func main() {
	//func expression , function can be assigned to variable
	f := func() {
		fmt.Println("fun expre")
	}
	f()
	fmt.Println(reflect.TypeOf(f)) //output : func()
}
