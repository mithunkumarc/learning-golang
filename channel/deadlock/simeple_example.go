package main

import (
	"sync"
)

func main() {
	// there is only one goroutine here: main goroutine
	m := sync.Mutex{}

	m.Lock() // main goroutine acquired lock from m
	// main goroutin again waiting to acquire lock from m
	m.Lock() // main goroutine wait forever and enter to deadlock
	m.Unlock()
	m.Unlock()
}

/*
  fatal error: all goroutines are asleep - deadlock!

*/
