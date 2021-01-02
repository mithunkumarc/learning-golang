package pointers
import "fmt"

/*
go support both pass by value and pass by ref
pass by ref 
*/


func PassByValueExample(){
  var x, y int = 3, 5
  fmt.Println("before calling swapping function")
  fmt.Println("x", x)
  fmt.Println("y", y)
  swap(x,y)
  fmt.Println("after calling swapping function")
  fmt.Println("x", x)
  fmt.Println("y", y) 
}


func swap(x , y int) {
  temp := x
  x = y
  y = temp
  fmt.Println("inside swapping function")
  fmt.Println("x ",x)
  fmt.Println("y ",y)
}


/*main.go

package main

import (
  "main/pointers"
)

func main() {
  pointers.PassByValueExample()
}
*/