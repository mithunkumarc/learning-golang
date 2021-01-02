package controlstatements

func PositiveOrNegetive(number int) string {
  if number < 0 {
    return "negetive"
  } else if number > 0 {
    return "positive"
  } else {
    return "zero"
  }
}


/* main.go

package main
import (
  "fmt"
  "main/controlstatements" // modulename/packagename
)

func main() {
 fmt.Println(controlstatements.PositiveOrNegetive(6)) // positive
 fmt.Println(controlstatements.PositiveOrNegetive(-6)) // negetive
 fmt.Println(controlstatements.PositiveOrNegetive(0)) // zero
}

*/