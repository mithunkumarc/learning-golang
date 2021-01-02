package datatypes
import (
  "fmt"
)
const Pi = 3.14
func Constants() {
  // Pi = 4.4; error
  const message = "hello world"
  fmt.Println(message)
}

/*main.go

package main
import (
  "main/datatypes" // modulename/packagename
  "fmt"
)
func main() {
  fmt.Println("Pi value : ", datatypes.Pi)
  datatypes.Constants()
}
*/