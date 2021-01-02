package structs
import (
"fmt"
)

type Pen struct {
  color string
}

func PassByReference() {
  var p1 Pen = Pen{color: "red"}
  var p2 Pen = Pen{color: "green"}
  // sent memory location of pen, 
  // if memory location is changes it affects both copy
  swapPen(&p1, &p2)
  fmt.Println("after calling swapping function")
  fmt.Println("p1 ",p1.color) // green
  fmt.Println("p2 ",p2.color) // red
}


func swapPen(p1 *Pen, p2 *Pen) {
  temp := p1.color
  p1.color = p2.color
  p2.color = temp
  fmt.Println("inside swapping function")
  fmt.Println("p1 ",p1.color) // green
  fmt.Println("p2 ",p2.color) // red
}
