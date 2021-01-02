package arrays
import (
  "fmt"
)
/*fixed number of elements with fixed type*/

func Introduction() {
//array declaration opposite to java
//java: int[] n = new int[5];
//golang: var n [5]int 
  var numbers [3]int
  var i int
  for i = 0 ; i < len(numbers); i++ {
    numbers[i] = i * 5
  }
  fmt.Println(numbers) // [ 0 5 10 ]
}

/* main.go
package main

import (
  "main/arrays"
  
)

func main() {
 arrays.Introduction()
}
*/