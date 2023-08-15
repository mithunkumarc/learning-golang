package main

import (
	"fmt"
)

func main() {
	// creating channel int
	dataChan := make(chan int) // unbuffered channel(no memory)
	dataChan <- 789
	n := <-dataChan
	fmt.Println(n)
}

/*
	output : 
	fatal error: all goroutines are asleep - deadlock
	Reason : By default channel(unbuffered) has no memory, 
	the moment we send data to channel we should read immediately.
	send data to channel and reading after it, will not work in 
	same goroutine(main thread)
*/
