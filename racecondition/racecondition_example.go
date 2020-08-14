package main

import (
	"fmt"
	"sync"
	"time"
)

type Account struct{
	name string
	balance int
}

var counter int = 0

// comment mux.Lock() and mux.Unlock() to see race conditions

func transaction(s *Account, d *Account, wg *sync.WaitGroup, mux *sync.Mutex, amount int) {
	defer wg.Done()	
	for i := 1 ; i <= 100 ; i++ {
		mux.Lock() // first goroutine gets lock, once current goroutine leaves other can get it 
		counter = amount // counter becomes the shared resource
		d.balance = d.balance + counter
		time.Sleep(1000) // allow another goroutine to create race condition
		s.balance = s.balance - counter
		counter = 0
		//fmt.Println(s,d)
		mux.Unlock() // release lock by current goroutine
	}
}

func main() {
	var wg sync.WaitGroup // required to make main routine to wait till all goroutine completes
	var mux sync.Mutex
	a := Account{name:"a", balance:100}
	b := Account{name:"b", balance:0}
	c := Account{name:"c", balance:200}
	d := Account{name:"d", balance:0}
	wg.Add(1)
	go transaction(&a, &b, &wg, &mux, 1)//incremented by 1
	//time.Sleep(10000000)
	wg.Add(1)
	go transaction(&c, &d, &wg, &mux, 2)//incremented by 2
	
	
	wg.Wait()
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
}
