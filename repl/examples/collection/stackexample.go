package collection

type Stack struct {
  // if you want to export this, make it Data
  data []string;
}

// func to add new element to stack
func (s *Stack) Push(x string) {
  // append is a global function
  s.data = append(s.data, x)
}

// func to remove element from stack
func (s *Stack) Pop() string {
  // get last index
  n := len(s.data) - 1
  res := s.data[n]
  s.data[n] = "" // avoid memory leak
  // resize stack, reduce len by 1
  s.data = s.data[:n]
  return res
}

func (s *Stack) Size() int {
  return len(s.data)
}

/*
source : https://yourbasic.org/golang/go-java-tutorial/#hello-stack-example

testing in main.go
package main
import (
"main/collection" // modulename/packagename
"fmt"
)
func main() {
  var s collection.Stack
  s.Push("world")
  s.Push("hello")
  fmt.Println(s.Size())
  fmt.Println(s)
}
*/