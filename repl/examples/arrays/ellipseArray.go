package arrays

import (
  "fmt"
)

// using ellipse , we can let compiler identify length
func EllipseArray() {
  words := [...]string{"apple","ball","cat"}
  fmt.Println("length of array : ",len(words))
  for _, word := range words {
    fmt.Println(word)
  }
}

/*main.go

package main

import (
  "main/arrays"
)

func main() {
 arrays.EllipseArray()
}
*/