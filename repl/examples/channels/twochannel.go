package channels

import (
  "fmt"
  "sync"
)

func add(ch1 chan int, ch2 chan int,wg *sync.WaitGroup) {
  defer wg.Done()//called when goroutine completed
  fmt.Println(<- ch1 + <- ch2)
}

func IntroductionAdd() {
  var wg sync.WaitGroup
  ch1 := make(chan int)
  ch2 := make(chan int)
  // send data to channel on separate goroutine
  // no need for waitgroup here
  go func(){ch1 <- 100}()
  go func(){ch2 <- 200}()
  //reading data from channel can be made in main routine
  wg.Add(1) //inform golang we adding one gorouting
  go add(ch1, ch2, &wg)  // adding in separate routine
  wg.Wait()// wait for my goroutine
  //closing channels
  fmt.Println("closing channels")
  close(ch1)
  close(ch2)
}

/*main.go

package main

import (
  "main/channels"
)

func main() {
 channels.IntroductionAdd()
}
*/