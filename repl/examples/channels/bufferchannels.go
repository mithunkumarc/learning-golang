package channels

import (
  "fmt"
)
/*buffered channels
data can be sent as block
mention the size of block while creating channel
*/

func acceptData(dc chan string) {
  fmt.Println(<- dc)
  fmt.Println(<- dc)
  fmt.Println(<- dc)
}


func BufferedChannelExample() {
  // only 2 string values are allowed
  dataChannel := make(chan string, 2)
  go func(){
  //if u are sending on goroutine. limit doesnt matter
  //below code doesn't complaint
  //exptected 2, but seding 3, still no error
    dataChannel <- "interstellar"
    dataChannel <- "inception"
    dataChannel <- "tenet"
  }()//self called function put on goroutine
  acceptData(dataChannel)
  // if are sending data on main routine then limit come to picture
   
}

/*main.go

package main

import (
  "main/channels"
)

func main() {
 channels.BufferedChannelExample()
}
*/