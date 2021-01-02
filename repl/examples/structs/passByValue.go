package structs
import (
"fmt"
)

type Balloon struct {
  color string
}

func PassByValue() {
  var b1 Balloon = Balloon{color: "red"}
  var b2 Balloon = Balloon{color: "green"}
  // sent ballon as xerox copy, 
  // if xerox copy modified doesnt change original copy
  swapBalloon(b1,b2)
  fmt.Println("after calling swapping function")
  fmt.Println("b1 ",b1.color) // still red
  fmt.Println("b2 ",b2.color) // still green
}


func swapBalloon(b1 Balloon, b2 Balloon) {
  temp := b1.color
  b1.color = b2.color
  b2.color = temp
  fmt.Println("inside swapping function")
  fmt.Println("b1 ",b1.color) // green
  fmt.Println("b2 ",b2.color) // red
}


/*main.go

package main

import (
  "main/structs"
  
)

func main() {
 structs.PassByValue()
}
*/