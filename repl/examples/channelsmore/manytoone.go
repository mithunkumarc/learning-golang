package channelsmore

import (
	"fmt"
)

/*
 	N-to-1 channel
	pushing data to same channel at multiple places
	reading at one places
*/

func ManyToOne() {
	c1 := make(chan int)  // source channel
	c2 := make(chan bool) // being used as WaitGroup: how? channel has capacity to wait till it receives the data.
	
	// sender 1 :d ata are sent on channels using goroutine except buffered channel 
	go func(){
		//source, pushing data to channel, 0 to 9
		for i := 0 ; i < 10 ; i++ {
			c1 <- i
		}
		c2 <- true // using c2 as waitgroup
		
		// i am not using close here as other places might be pushing data to this channel
	}()
	
	// sender 2 : sending data on channel c1
	go func(){
		//source, pushing data to channel, 10 to 19
		for i := 10 ; i < 20 ; i++ {
			c1 <- i
		}
		c2 <- true //using c2 as waitgroup
		//not using close(c1) here as we don't know whether other place is done with pushing data or not
	}()
	
	// only one receiver  : reading the data fromt the channel
	go func(){
		for i := 0 ; i < 20 ; i++ {
			fmt.Println(<- c1)
		}
		/*
		using c2 as WaitGroup, now anywhere if u are using "<- c2", 
		program will wait till data is received 
		*/ 
		c2 <- true	
	}()
	
	
	<- c2	// c2 is waiting till it receives data, so main goroutine still running
	<- c2	// c2 is waiting till it receives data, so main goroutine still running
	<- c2
	close(c1) // safe to close here
	// don't expect in sequnce order. gorouine/thread are unpredictable
	// using channel as waitgroup is bad practice
}

/*main.go


package main

import (
  "main/channelsmore"
)

func main() {
 channelsmore.ManyToOne()
}
*/