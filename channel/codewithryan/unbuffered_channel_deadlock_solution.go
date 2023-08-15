package main

import (
	"fmt"
)

func main() {
	// creating channel int in a separate goroutine
	dataChan := make(chan int) // unbuffered channel(no memory)
	go func() {
		dataChan <- 789
	}()
	n := <-dataChan
	fmt.Println(n)
}

/*
	output : 789
	unbuffered deadlock solution :
	in one goroutine(thread) : send data
	and in another goroutine(thread) receive data
 	--or use-- buffered channel
*/
