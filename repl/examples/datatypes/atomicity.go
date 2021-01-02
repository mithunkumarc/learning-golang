package datatypes

/*

java has AtomicInteger : 
it can be used as counter in multithreaded environment.
becuase only one thread can access at a time.
in the same way atomic types are used in golang
in multithreaded environment.
atomic types can be accessed only one routine at a time.
if you dont want to use atomic in golang. use sync.Mutex Lock and Unlock function to handle counter 
manually.

eg for atomic variables : int32, int64, uint64
find all in https://golang.org/pkg/sync/atomic/

If you use normal dataype in multithreaded as counter, 
you may expect wrong value at the end of program
*/


import (
    "fmt"
    "sync"
    "sync/atomic"
)

// alternate to this u can use Mutex Lock/Unlock
var ops uint64 // counter var, thread safe atomic type
//if you use non atomic type, wrong counter value

func task(wg *sync.WaitGroup) {
	for i := 0 ; i < 10 ; i++ {		
		atomic.AddUint64(&ops, 1)//one thread at a time can access
	}
	wg.Done()
}


func AtomicExample() {

    
    var wg sync.WaitGroup

    for i := 0; i < 50; i++ {
        wg.Add(1)
	go task(&wg) // 50 times * 10 times = output ops counter must be 500
    }

    wg.Wait()//wait all goroutine to finish

    fmt.Println("ops:", ops) // 500
    //dont expect counter value sequentially
    //but right value output is guaranteed
}


/*main.go

package main

import (
  "main/datatypes"
)

func main() {
 datatypes.AtomicExample()
}

*/