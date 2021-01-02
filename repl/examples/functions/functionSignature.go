package functions
import "fmt"

type someOperation func(a,b int) int

// follows someOperation signature
func add(a,b int) int {
  return a + b
}

func FunctionSignatureExample() {
  var myAdd = add // assinging to another var
  res := myAdd(3,6)
  fmt.Println(res)
}

/*main.go


package main

import (
  "main/functions"
)

func main() {
 functions.FunctionSignatureExample()
}
*/