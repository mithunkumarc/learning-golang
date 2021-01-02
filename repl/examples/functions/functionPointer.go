package functions
import (
  "fmt"
)

/*
No functional pointer in go, you can create 
function signature
*/
type hello func(string) string

func FunctionPointerExample() {
  msg := helloAcceptor(helloExample)
  fmt.Println(msg)
}

// signature acceptor, any function with hello type accepted
func helloAcceptor(h hello) string {
  return h("bharat")
}

// signature implementation
func helloExample(name string) string{
  return "hello, "+name
}