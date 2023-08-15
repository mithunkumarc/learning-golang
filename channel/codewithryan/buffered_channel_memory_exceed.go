package main

import (
	"fmt"
)

func main() {
	dataChan := make(chan int, 1) // buffered channel(has memory size 1)
	dataChan <- 789
	dataChan <- 121 // deadlock , no space for second data
	n := <-dataChan
	fmt.Println(n)
}

/*
	output : 
	fatal error: all goroutines are asleep - deadlock!
	memory size is 1, but sending two data
 	solution : increasd the memory size for channel from 1 to 2
*/
