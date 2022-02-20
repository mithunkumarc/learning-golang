package main

import (
	"fmt"
)

func main() {
	s := "hello"
	c := 'J' // no char type in golang, but it uses int32
  // Printf() used to print type of variable
  // format specifier %T used to print type of variable
	fmt.Printf("type of %c is %T \n",c,c) // type of J is int32
	fmt.Printf("%T \n",s) // string
}

// https://www.callicoder.com/golang-basic-types-operators-type-conversion/
