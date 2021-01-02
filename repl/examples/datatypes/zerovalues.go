package datatypes
import "fmt"

/*zero values nothing but default/empty value of datatype*/

func ZeroValues() {
  var i int
  var b bool
  var s string
  fmt.Println("default value of int is ", i) // 0
  fmt.Println("default value of bool is ", b) // false
  fmt.Println("default value of string is ", s) // empty string
}

/* test this file from main.go

package main
import (
  "main/datatypes"
)
func main() {
  datatypes.ZeroValues()
}
*/