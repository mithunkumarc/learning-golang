package controlstatements
import (
  "fmt"
)
//if both parameters of same type, declare together
func PrintEvenNumbers(start , end int) {
  fmt.Printf("printing even numbers between %d and %d \n", start , end)
  for i := start; i <= end ; i++ {
    if i % 2 == 0 {
      fmt.Println(i)
    }
  }
  fmt.Println("printing numbers in reverse order")
  for j := end; j >=start ; j-- {
    fmt.Println(j)
  }
}


/*main.go

package main
import (
 "main/controlstatements"
)

func main() {
 controlstatements.PrintEvenNumbers(1, 10) 
}
*/