declaring package and import package

        package main
        import (
          "fmt"
        )
        
declaring struct

        // if you are not exporting outside package, smallercase is fine
        type Tiger struct {
          name string
        }

link method Eat to Tiger struct

        // if you are not exporting function outside package smallercase if fine
        func (t Tiger) Eat() {
          fmt.Println("tiger eating")
        }

main function 

        func main() {
          var t Tiger = Tiger{name: "huliraja"}
          t.Eat()
        }
