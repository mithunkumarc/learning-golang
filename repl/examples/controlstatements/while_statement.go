package controlstatements
import (
  "fmt"
)
/*
  there is no while statemet in go
  for loop can be used as while loop
*/
func PrintNumbers(breakPoint int) {
  i := 0
  for ; i < breakPoint;  {
    i++
    fmt.Println(i)
  }
}

/*main.go

package main
import (
 "main/controlstatements"
)

func main() {
 controlstatements.PrintNumbers(10) 
}
*/