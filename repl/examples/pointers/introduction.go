package pointers

import (
  "fmt"
)

/*
  similar to c/c++
  *p declaring pointer
  &p reading address
  *p deferencing
*/

func Introduction() {
  i := 32
  j := &i // j is a integer pointer, type inferred
  fmt.Println(i) // 32
  fmt.Println(j) // prints address
  fmt.Printf("type of j is %T \n", j) // *int
  fmt.Println("deferencing ", *j) // 32
}

/*main.go

package main

import (
  "main/pointers"
)

func main() {
  pointers.Introduction()
}
*/