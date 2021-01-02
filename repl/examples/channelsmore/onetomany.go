package channelsmore

import (
	"fmt"
)

/*
 	1-to-N channel
	one source channel
	multiple receivers//reading at many places
*/

func OneToMany() {
	c1 := make(chan int)  // source channel
	c2 := make(chan bool) // being used as WaitGroup: how? channel has capacity to wait till it receives the data.
	
	// data are sent on channels using goroutine except buffered channel 
	go func(){
		//source, pushing data to channel
		for i := 0 ; i < 10 ; i++ {
			c1 <- i
		}
		close(c1) // good to  close channel after you are done
	}()
	
	// receiver 1 : reading half of the data in channel
	go func(){
		for i := 0 ; i < 5 ; i++ {
			fmt.Println(<- c1, "hero")
		}
		/*
		using c2 as WaitGroup, now anywhere if u are using "<- c2", 
		program will wait till data is received 
		*/ 
		c2 <- true	
	}()
	
	// receiver 2 : reading half of the data in channel
	go func(){
		for i := 0 ; i < 5 ; i++ {
			fmt.Println(<- c1, "heroine")
		}
		/* using c2 as waitGroup */
		c2 <- true
	}()
	<- c2	// c2 is waiting till it receives data, so main goroutine still running
	<- c2	// c2 is waiting till it receives data, so main goroutine still running
	// don't expect in sequnce order. gorouine/thread are unpredictable
  // using channel as waitgroup is bad practice
}


/*main.go

package main

import (
  "main/channelsmore"
)

func main() {
 channelsmore.OneToMany()
}
*/