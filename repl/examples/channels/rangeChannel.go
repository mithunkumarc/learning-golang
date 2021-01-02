package channels

import ("fmt")

/*
range can be used to get/read elements from  channel.
to use range, channel should be closed.
else error will be fired at the end.
*/

func RangeChannelExample() {
  numChannel := make(chan int)
  go func() {
    numChannel <- 2
    numChannel <- 3
    numChannel <- 4
    close(numChannel)
  }()
  
  //this will not work if numChannel is closed in same routine as you are reading. it must be different
  for i := range numChannel {
    fmt.Println(i)
  }
}

/*main.go

package main

import (
  "main/channels"
)

func main() {
 channels.RangeChannelExample()
}
*/