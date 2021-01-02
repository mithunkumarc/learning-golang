package controlstatements

func OddOrEven(number int) string {
  if number % 2 == 0 {
    return "even"
  } 
  return "odd"
}


/*main.go

package main
import (
  "fmt"
  "main/controlstatements"
)

func main() {
  fmt.Println(controlstatements.OddOrEven(4))
  fmt.Println(controlstatements.OddOrEven(41))  
}
*/