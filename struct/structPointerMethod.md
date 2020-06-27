pass by reference : object state will be change after operation

declaring and importing package

        package main
        import (
          "fmt"
        )

declaring struct

      // if you are not exporting outside package, smallercase is fine
      type Tiger struct {
        name string
      }

change name of tiger, argument is sent as pointer

        //receives address of object
        func (t *Tiger) ChangeName(){
          t.name = "hulianna";    // original state of tiger is permanently changed
        }

main function 

        func main() {
          var t Tiger = Tiger{name: "huliraja"}
          fmt.Println("before change", t.name); // huliraja
          t.ChangeName();   // sending reference (address)
          fmt.Println("after change", t.name); // hulianna
        }

NOTE: if method received argument using pass by value, object state wil not be changed permanently
        
        // not a pointer argument
        func (t Tiger) ChangeName(){
          t.name = "hulianna";
        }
        
