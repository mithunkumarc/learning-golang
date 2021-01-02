package channels
/*
  name itself explains
  datatype is fixed
  send or receive data
  goroutine can communicate each other using channel 
  <- operator used to show direction of data going
*/
import (
  "fmt"
)
/*
creating channel
 use  : make(chan DataType)
 using go routine to send data on channel 
 why? 
*/

func Introduction() {
  //creating channel
  var mychannel chan int = make(chan int)
  fmt.Printf("datatype :%T \n", mychannel)// chan int
  //end data to chan in separate goroutine
  go func() {mychannel <- 34}()
  fmt.Println(mychannel)// prints address
  //main routine reading data from channel
  data := <- mychannel 
  //anotherData := <- mychannel //error : gets data one time only
  fmt.Println(data)
}