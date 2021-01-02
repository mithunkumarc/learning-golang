package datatypes

import (
	"fmt"
	"unsafe"
)

func DatatypeSize() {
	var i int
	var u int32
	var g int64
  //depends on os bit
  fmt.Printf("i Type:%T Size:%d\n", i, unsafe.Sizeof(i)) 
  fmt.Printf("u Type:%T Size:%d\n", u, unsafe.Sizeof(u)) // 4bytes
  fmt.Printf("g Type:%T Size:%d\n", g, unsafe.Sizeof(g)) // 8bytes
}
/*
test in main.go

package main
import (
  "main/datatypes"
)
func main() {
  datatypes.DatatypeSize()
}
*/