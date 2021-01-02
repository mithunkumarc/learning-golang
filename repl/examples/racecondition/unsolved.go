package racecondition

import (
  "fmt"
  "time"
)


var counter int = 0

func task(m string, n int) {
  for i := 0 ; i < 100 ; i++ {
    counter = n
    time.Sleep(100000)
    fmt.Println(m, counter)
    counter = 0
  } 
}

func RaceConditionExample() {
  go task("first", 1) // first should get 1
  go task("second", 2) // second should get 2
  time.Sleep(time.Second)
}

/*main.go

package main

import (
  "main/racecondition"
)

func main() {
 racecondition.RaceConditionExample()
}

output : 
second 1
second 2
first 2
first 1
...
*/