package controlstatements
import "fmt"

func ForRangeExample() {
  numbers := [3]int{1,2,3}
  for index, number := range numbers {
    fmt.Println(index, ":", number)
  }
}

/*main.go

package main

import (
  "main/controlstatements"
)

func main() {
 controlstatements.ForRangeExample()
}
*/