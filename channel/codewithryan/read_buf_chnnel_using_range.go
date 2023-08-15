package main

import (
	"fmt"
	"time"
)

func main() {
	dataChan := make(chan int, 2) // buffered channel(has memory)

	dataChan <- 789
	dataChan <- 121
	close(dataChan)             // if you miss this, deadlock error thrown
	time.Sleep(2 * time.Second) // let channel close
	for d := range dataChan {
		fmt.Println(d)
	}
}

/*
	output :
	789
	121
*/
