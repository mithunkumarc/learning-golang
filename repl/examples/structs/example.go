package structs

import (
  "fmt"
)

//visible to outside package
type Vertex struct {
  x, y int // not accessible outside as they begin with lowercase
}

/*goang def :
struct is a collection of fields
*/
func StructExample() {
  var v Vertex = Vertex{10,-43}
  fmt.Println(v.x)
  fmt.Println(v.y)
  fmt.Println(v) // {10 -43}
}

/*main.go

package main

import (
  "main/structs"
)

func main() {
 structs.StructExample()
}
*/