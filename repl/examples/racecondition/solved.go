package racecondition

/*
check this program often.
especially snc.Mutex

sync.Mutex uses two important methods to synchronize
Lock() similar to java thread synchronize lock
Unlock() release lock
*/

import (
    "fmt"
    "time"
    "sync"
)

var counter1 int = 0

func task1(m string, n int,  wg *sync.WaitGroup, mux *sync.Mutex) {
    
    defer wg.Done()
    for i := 0 ; i < 100 ; i++ {
        mux.Lock() // synchronizing gorouting, one at a time
        //your shared resource
        counter1 = n
        time.Sleep(100000)
        fmt.Println(m, counter1, i)
        counter1 = 0
        //end of your shared resource
        mux.Unlock() // end of synchronizing, other routine can use
    }
}

func SyncMutexRaceCondition() {
    var wg sync.WaitGroup
    var mux sync.Mutex
    wg.Add(1) //inform main routine
    go task1("first", 1, &wg, &mux)// first should get 1
    wg.Add(1)
    go task1("second", 2, &wg, &mux)// second should get 2
    wg.Wait()//wait till all goroutine completes
}

/*main.go

package main

import (
  "main/racecondition"
)

func main() {
 racecondition.SyncMutexRaceCondition()
}

*/