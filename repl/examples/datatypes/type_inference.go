package datatypes
import "fmt"

/*
  similar to java type inference
  in java : 
  var i = 10; compiler assign type int
  i = false; error as type already assigned as int

*/
func TypeInference() {
  i := 10 // inferred type is int
  fmt.Printf("%T \n ", i) // int
  // i := true //error
}