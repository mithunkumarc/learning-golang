package controlstatements
import (
  "fmt"
)
/*switch without condition(option) 
same as switch true {} : switch bool_value {}
*/
func SwitchNoCondition() {
  option := 5
  switch {
    case option < 5: fmt.Println("less than 5")
    case option > 5: fmt.Println("more than 5")
    default: fmt.Println("equal to 5")
  }
}

/*main.go

package main
import (
 "main/controlstatements"
)

func main() {
  controlstatements.SwitchNoCondition()
}
*/