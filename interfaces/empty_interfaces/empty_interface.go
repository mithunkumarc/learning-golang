/*
  The interface type that specifies zero methods is known as the empty interface:

  interface{}
  An empty interface may hold values of any type. (Every type implements at least zero methods.)

  Empty interfaces are used by code that handles values of unknown type. 
  For example, fmt.Print takes any number of arguments of type interface{}

*/

package main

import (
	"fmt"
)

func main() {
	var i interface{} = "hello world"
	string, flag := i.(string)
	if flag {
		fmt.Println(string)
	} else {
		fmt.Println("not a string value")
	}

}
