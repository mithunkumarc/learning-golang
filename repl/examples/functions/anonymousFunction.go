package functions
import "fmt"

func AnonymousFunctionExample() {
  adder := func(a,b int)int{return a + b}
  fmt.Println(adder(3,5))
}

/* main.go

package main

import (
  "main/functions"
)

func main() {
 functions.AnonymousFunctionExample()
}
*/