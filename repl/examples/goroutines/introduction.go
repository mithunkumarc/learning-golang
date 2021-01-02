package goroutines
import (
  "fmt"
  "time"
)
/*
two goroutines
1. main goroutine : default thread like java
2. user goroutine : we can create our own goroutine
*/

func GoRoutineExample() {
  go task1() // separete go routing : parallel executiong
  /*if you put task2() into another goroutine like 
  go task2(), another routine will be created but main 
  routing will not wait for task1() and task2(). program exits: solution : see (0)
  */
  task2() // executed by main goroutine 
  // (0) : time.sleep(enough to complete two routines)
  // not a good practice though
}
func task1() {
  for i := 1 ; i < 10 ;i++ {
    fmt.Println("task1")
    time.Sleep(1000000000)
  } 
}
func task2() {
  for i := 1 ; i < 10 ;i++ {
    fmt.Println("task2")
    time.Sleep(1000000000)
  } 
}

/*main

package main

import (
  "main/goroutines"
)

func main() {
 goroutines.GoRoutineExample()
}
*/