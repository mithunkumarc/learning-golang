package main

import (
	"fmt"
)

func main() {
	dataChan := make(chan int) // unbuffered channel(no memory)
	go func() {
		dataChan <- 789
		dataChan <- 121
		close(dataChan) // if you miss this, deadlock error thrown
	}()
	for d := range dataChan {
		fmt.Println(d)
	}
}

/*
	output :
	789
	121
	using separate goroutine because of unbuffered channel
	data is read uging range
*/
