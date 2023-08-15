package main

import (
	"fmt"
	"time"
)

var (
	count int
)

func main() {
	iterations := 1000
	for i := 0; i < iterations; i++ {
		go increment()  // multiple goroutine trying increment count variable simultaneously
	}
	time.Sleep(5 * time.Second)
	fmt.Println("count : ", count)
}

func increment() {
	count++
}


/*
  output : 
  expected : 1000
  actual : less than 1000
*/
