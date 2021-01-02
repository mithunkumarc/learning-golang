package goroutines

import (
  "fmt"
  "sync"
  "time"
)

/*if main goroutine complete its task
it will not wait other routine to complete
solution : waitgroup
waitgroup only waits for all goroutines to complete.
important : waitgroup cannot be used for synchronization.
sync.Mutex is used to synchronize goroutine.

*/

var globalCounter = 0;

func IntroductionExample() {
  var wg sync.WaitGroup
  /*incrementing wg counter by 1. because we are adding go routines. or you can declare altogether
  wg.Add(numberOfroutineYOuwantToAdd)
  */
  wg.Add(1)
  go printMessage("hello", &wg, 1)
  //if above call complete quickly : (0)
  wg.Add(1)
  go printMessage("world", &wg, 2)
  //add time.Sleep to see go printMessage output
  wg.Wait()
}

func printMessage(msg string, wg *sync.WaitGroup, j int) {
  defer wg.Done()// decrements wg counter by 1
  for i := 0 ; i < 10 ; i++ {
    globalCounter =  j
    fmt.Println(msg, globalCounter)
    time.Sleep(time.Second)
  }
  //i guess you can put wg.Done() here without defer
}
