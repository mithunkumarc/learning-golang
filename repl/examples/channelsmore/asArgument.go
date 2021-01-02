package channelsmore

import "fmt"

func pushData() chan int {
	c := make(chan int)
	go func() {
		for i := 1; i <= 10 ; i++ {
			c <- i
		}
		close(c)	
	}()
	return c
}

func receiveData(c chan int) {
	for data := range c {
		fmt.Println(data)
	}
}

func AsArgument() {
	c := pushData()
	receiveData(c) // sending channel as argument
}


/*main.go

package main

import (
  "main/channelsmore"
)

func main() {
 channelsmore.AsArgument()
}

*/