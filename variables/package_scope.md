#### variable can be declared outside function, this scope is package scope

      example:
      
      package main
      import "fmt"
      
      var color := "green"    // package scope variable

      func main() {
        //..
      }
