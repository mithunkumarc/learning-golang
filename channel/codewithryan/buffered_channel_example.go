package main

import (
	"fmt"
)

func main() {
	dataChan := make(chan int, 1) // buffered channel(has memory)
	dataChan <- 789
	n := <-dataChan
	fmt.Println(n)
}

/*
	output : 789
   sending to and reading from channel in the same goroutine works
   because channel has memory
*/
