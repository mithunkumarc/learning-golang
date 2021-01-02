package interfaces

// exporting rectange struct (upper case)
type Rectange struct {
  /*
   properties also exported
   if u dont want this, then send len breadth 
   from outside and create struct Rectange in this file
  */
  Length, Breadth int
}

/* 
contract : any struct which implements Area() then 
it can be treated as Shape
*/
type Shape interface {
  // exporting
  Area() int
}

func (r Rectange) Area() int {
  return r.Length * r.Breadth
}

// calcualtes area, send struct which implemnts area
func CalculateArea(s Shape) int {
  return s.Area()
}

/* //code inside main.go to test
package main
import (
  "main/interfaces" // modulename/packagename
  "fmt"
)
func main() {
  r := interfaces.Rectange{10, 10};
  area := interfaces.CalculateArea(r)
  fmt.Println(area)
}

*/