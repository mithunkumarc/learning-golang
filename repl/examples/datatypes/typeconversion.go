package datatypes

import (
  "fmt"
)

func TypeConversion() {
  // type inferred is int same as int64 in 64bit os
  i := 10 
  fmt.Printf("datatype of i is %T \n", i)// int
  // int to int64
  j := int64(i)
  fmt.Printf("datatype of j is : %T \n", j) // int64
  // float64 to float32
  k := float64(i)
  m := float32(k)
  fmt.Printf("datatype of m is : %T \n", m) // float32
}

