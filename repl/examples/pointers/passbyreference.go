package pointers

import (
  "fmt"
)

func PassByReferenceExample() {
  var x, y int = 2, 4
  fmt.Println("before swapping")
  fmt.Println("x ", x)
  fmt.Println("y ", y)
  //swap already declared in other file same package
  swapByReference(&x, &y)
  fmt.Println("after swapping")
  fmt.Println("x ", x)
  fmt.Println("y ", y)
}

func swapByReference(x , y *int) {
  temp := *x
  *x = *y
  *y = temp
  fmt.Println("inside function swap ")
  fmt.Println("x ", *x)
  fmt.Println("y ", *y)
}


/*main.go

package main

import (
  "main/pointers"
)

func main() {
  pointers.PassByReferenceExample()
}
*/