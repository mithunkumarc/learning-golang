package arrays
import (
  "fmt"
)
func StringArrayExample() {
  var words [5]string;
  words[0] = "monday"
  words[1] = "tuesday"
  words[2] = "wedday"
  words[3] = "thursday"
  words[4] = "friday"

  for i := 0; i < len(words) ; i++ {
    fmt.Println(words[i])
  }
  // range returns index and element 
  // _ : i don't want to use index
  // if you declare var for index and didnt use it, gives error
  for _ , content := range words {
    fmt.Println(content)
  }
}

/*main.go

package main

import (
  "main/arrays"
  
)

func main() {
 arrays.StringArrayExample()
}
*/

