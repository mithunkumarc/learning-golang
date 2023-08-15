package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	count int
	lock  sync.Mutex
)

func main() {
	iterations := 1000
	for i := 0; i < iterations; i++ {
		go increment()
	}
	time.Sleep(5 * time.Second) // wait till all goroutine complete
	fmt.Println("count : ", count)
}

func increment() {
	lock.Lock()   //  goroutine getting lock
	count++
	lock.Unlock() // gorouting return lock
}


/*
  output : 
  expected : 1000
  actual : 1000
*/
