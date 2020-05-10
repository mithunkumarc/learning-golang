package main

import (
	"fmt"
)

func main() {
  //similar to self calling functions in javascript, python
	func(x int) {
		fmt.Println("anony",x)
	}(42)	
}

//anony 42
//in python : c = (lambda x: x * x)(5)
