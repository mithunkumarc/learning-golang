package structs
import (
  "fmt"
)
/*
struct literal is nothing but when u print 
stuct you will find list of values enclosed in curly braces. 
*/

type Mobile struct {
  color string
  cost float64
}
func StructLiteral() {
  var (
    m = Mobile{"red",232}
    m1 = Mobile{color:"white"} // cost will be zero
    m2 = Mobile{cost:232} // color will be empty string
    m3 = Mobile{} // empty string and zero
    m4 = &Mobile{"green",232}
  )
  fmt.Println(m) // struct literal : {red 232}
  fmt.Println(m1) // struct literal : {white 0}
  fmt.Println(m2) // struct literal : {232}
  fmt.Println(m3) // struct literal : {0} 
  fmt.Println(m4) // struct literal : &{green 232}
}

/*main.go

package main

import (
  "main/structs"
  
)

func main() {
 structs.StructLiteral()
}
*/