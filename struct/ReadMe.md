A struct is collection of fields. similar in c/c++.
methods defined can be linked to struct.

        package main

        import (
          "fmt"
        )

        type Person struct {
          name string
        }
        func (p Person) speak() {
          fmt.Println(p.name,"speaks")
        }

        func main() {
          p := Person{name:"raj"}
          p.speak()	
        }
