package main

import (
	"fmt"
)

// var x int = 0 // wrong idea

func main() {
	//closure : containing the var inside a block
	a := incrementor() // seperate memory
	b := incrementor() // seperate memory
	fmt.Println(a()) // 1
	fmt.Println(a()) // 2
	fmt.Println(a()) // 3
	fmt.Println(a()) // 4
	fmt.Println(b()) // 1
	fmt.Println(b()) // 2
}

func incrementor() func() int {
  // scope of variable x is limited only in x , not accessible outside to change
	var x int = 0
	return func () int {
		x++
		return x
	}
}
