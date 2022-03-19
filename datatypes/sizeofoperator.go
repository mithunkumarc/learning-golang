package main

import (
	"fmt"
	"unsafe"
)

func main() {
	c := 100
	var d int8 = 20
	fmt.Printf("%T %T \n",c,d) // int int8
  fmt.Printf("%d bytes\n", unsafe.Sizeof(c)) // int 8 bytes, size of int is 8 bytes
}
