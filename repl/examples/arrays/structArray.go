package arrays

import (
  "fmt"
)

type Bus struct {
  color string
  numberOfSeats int
}

func StructArrayExample() {
  bus1 := Bus{"red",32}
  bus2 := Bus{"green",22}
  bus3 := Bus{"white",20}
  buses := [...]Bus{bus1, bus2, bus3}
  for _, bus := range buses {
    fmt.Println(bus)
  }
}

/*main.go
package main

import (
  "main/arrays"
)

func main() {
 arrays.StructArrayExample()
}
*/