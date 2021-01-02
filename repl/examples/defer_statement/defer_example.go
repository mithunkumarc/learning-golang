package defer_statement
/*
i made mistake in this propgram
lesson learnt : 
package name must match enclosing folder of this file.
*/

/*
defer statements executed last 
non defer statements executed first
defer statements goes to stack. 
defer executed in LIFO manner
*/

import (
  "fmt"
)

func DeferExample() {
  defer task1() // go to stack
  task2() // executes first
  defer task3() // go to stack 
  // in stack task3() is on top
  // task3() executes first 
  // task1() executes last
}

func task1() {
  fmt.Println("task1")
}

func task2() {
  fmt.Println("task2")
}

func task3() {
  fmt.Println("task3")
}


/*main.go

package main

import (
  "main/defer_statement"
)

func main() {
  defer_statement.DeferExample()
}
*/