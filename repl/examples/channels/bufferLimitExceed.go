package channels
import "fmt"

func BufferChannelLimit() {
  buffChannel := make(chan int, 2)
  buffChannel <- 100
  buffChannel <- 200
  //buffChannel <- 300 //error , limit exceeded
  fmt.Println(<- buffChannel)
  fmt.Println(<- buffChannel)
}


/*main.go

package main

import (
  "main/channels"
)

func main() {
 channels.BufferChannelLimit()
}
*/