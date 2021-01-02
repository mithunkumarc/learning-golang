package functions
import "fmt"

/*
golang executes init function first before executing
main().
you can have any number of init functions.
*/

func init() {
  fmt.Println("this is init method")
}

func init() {
  fmt.Println("this is another init method")
}

func InitExample() {
  fmt.Println("hello world")
}

/*main.go

package main

import (
  "main/functions"
)

func main() {
 functions.InitExample()
}

*/