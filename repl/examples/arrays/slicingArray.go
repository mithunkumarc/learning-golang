package arrays
import (
  "fmt"
)

func SlicingArrayExample() {
  var numbers = [...]int{2,4,6,8,10,12}
  fmt.Println("all elements")
  fmt.Println(numbers)
  // selecting {4,6,8}
  fmt.Println("selecting elements from index 1 to 3")
  fmt.Println(numbers[1:4])
  //if u dont provide end index,  then it selects till the end of array
  fmt.Println("skipping end index [2:]")
  fmt.Println(numbers[2:])
  //if u dont provide start index, it will being from index 0
  fmt.Println("skipping start index [:4]")
  fmt.Println(numbers[:4])

} 

/*main.go 
package main

import (
  "main/arrays"
)

func main() {
 arrays.SlicingArrayExample()
}
*/