package main

import (
	"fmt"
	"time"
)

func main() {
	one := make(chan string)
	two := make(chan string)
	quit := make(chan string)

	go func() {
		time.Sleep(1 * time.Second)
		one <- "Hey"
	}()
	go func() {
		time.Sleep(1 * time.Second)
		two <- "Hello"
	}()
	go func() {
		time.Sleep(1 * time.Second)
		quit <- "quit"
	}()

	for {
		select {
		case rec1 := <-one:
			fmt.Println("rec1 from one channel, msg = ", rec1)
		case rec2 := <-two:
			fmt.Println("rec2 from two channel, msg = ", rec2)
		case q := <-quit:
			fmt.Println("terminating message from channel = ", q)
			return
		}
	}
}

/*
  without for loop : select works with only one case
  use for loop to cover multiple case
  use for index based finite loop or dedicated quit channel to quit for-select

*/

