package controlstatements

import (
  "fmt"
)

func SwitchString(day string) {
  switch day {
    case "monday" : printData(day)
    case "tuesday" : printData(day)
    case "wednesday" : printData(day)
    case "thursday" : printData(day)
    case "friday" : printData(day)
    default: fmt.Println("weekend: Don't Distrub")
  }

}

// interanl caluculation, not exported
func printData(day string) {
  fmt.Println("today is ",day)
}

/*main.go

package main
import (
 "main/controlstatements"
)

func main() {
 controlstatements.SwitchString("friday")
 controlstatements.SwitchString("sunday") 
}
*/