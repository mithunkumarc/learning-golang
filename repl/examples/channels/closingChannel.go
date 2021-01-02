package channels
import "fmt"

/*
  if you close channel , you cannot send/read data anymore.

  while reading channel it returns two values.
  one value is data and another is has Channel contains more data.

*/

func readChannel(c1 chan int) {
  data, hasMoreData := <- c1
  fmt.Println(data, hasMoreData)// 400, true
  data, hasMoreData = <- c1
  fmt.Println(data, hasMoreData)// 100, true
  data, hasMoreData = <- c1
  fmt.Println(data, hasMoreData)//200, true
  data, hasMoreData = <- c1
  fmt.Println(data, hasMoreData)//300,true
  //data, hasMoreData = <- c1 // error, no more data
  //its better to close channel here
  close(c1)//no more data send/received on channel c1
}

func ClosingChannel() {
  c1 := make(chan int)
  go func(){
    c1 <- 100
    c1 <- 200
    c1 <- 300
  }()
  go func(){
    c1 <- 400 // works because channel is not closed yet
  }()
  readChannel(c1)
}


/*main.go


package main

import (
  "main/channels"
)

func main() {
 channels.ClosingChannel()
}
*/