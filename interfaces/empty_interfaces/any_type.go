package main

/*
  any type alias for empty interface
*/

import (
	"fmt"
)

func main() {
	var i any = "hello world"
	string, flag := i.(string)
	if flag {
		fmt.Println(string)
	} else {
		fmt.Println("not a string value")
	}

}
