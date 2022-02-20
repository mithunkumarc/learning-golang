#### variables declared in package scope can be exported to other packages

examples:

##### demo.go

      package hello
      var color = "green"

##### main.go

      package main
      import (
        "github.com/mithunkumarc/mygo/hello"
        "fmt"
      )

      func main() {
        fmt.Println(hello.color)  // undefined error
      }

#### export package scope variables by declaring with Captial case(beginning letter in upper case)

##### demo.go

      package hello
      var Color = "green"
      
##### main.go

      package main
      import (
        "github.com/mithunkumarc/mygo/hello"
        "fmt"
      )

      func main() {
        fmt.Println(hello.Color)  // green
      }
