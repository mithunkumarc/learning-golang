// waitgroup not include

#### generics with comparable

        package main

        import "fmt"

        /*
        The new comparable keyword, in Go 1.18, was added for specifying types that can be compared with the == and != operators.
        Comparable types include: structs, pointers, interfaces, channels, and builtin types. 
        */
        func displayData[T comparable](x T) {
          fmt.Printf("the type os x is %T \n", x)
        }

        func main() {
          displayData(1)
          displayData("hello world")
        }
        
#### singly linked list

        package main

        import "fmt"

        type Node struct {
          value int
          next  *Node
        }

        func NewNode(value int) *Node {
          var n Node
          n.value = value
          n.next = nil
          return &n
        }
        func main() {
          head := NewNode(30)
          head.next = NewNode(10)
          fmt.Println(head.value)
          fmt.Println(head.next.value)
        }

#### channel used for communication between goroutines

        package main

        import (
          "fmt"
        )

        func sum(s []int, c chan int) {
          sum := 0
          for _, value := range s {
            sum = sum + value
          }
          c <- sum // for sending one value, chan may not complain
        }

        func main() {
          var s []int = []int{1, 2, 4, 5}
          var c chan int = make(chan int)
          go sum(s, c)
          x := <-c 	//wait for result
          fmt.Println(x)
        }


#### buffered channel : storage capcity of channel

        package main

        import (
          "fmt"
        )

        /*
        Buffered channel

         Channels can be buffered.
        Provide the buffer length as the second argument to make to initialize a buffered channel:

        ch := make(chan int, 100)
        if you exceed the capacity of buffer you will get error
        channel takes value only when it has empy space
        */

        func main() {
          var ch chan int = make(chan int, 2)
          ch <- 1
          ch <- 2
          //ch <- 3 //error , no space left in channel
          fmt.Println(<-ch) // prints 1 (1 removed/read from channel)
          fmt.Println(<-ch) // prints 2
          //now channel is empty now it is ready to receive data
          ch <- 3 // works
        }

#### If you use range to read channel then use close in sender so that to indicate range that no more messages left

        package main

        import (
          "fmt"
        )

        /*
        range
        close

        if receiver is reading channel using range, then use close 
        in the sender indicating channel is close so no more data is coming

          v, ok := <-ch      // check any data left
          ok : false for no more data and true means still data left
        */
        func Sender(ch chan int) {
          var data []int = []int{100, 200, 300, 400}
          for _, value := range data {
            ch <- value
          }
          close(ch)
        }
        func main() {
          var ch chan int = make(chan int, 4)
          Sender(ch)
          for value := range ch {
            fmt.Println(value)
          }
          v, ok := <-ch      // check any data left
          fmt.Println(v, ok) // if no then ok prints false, v is 0
        }

#### golang select : can be used to work with multiple channels 


        package main

        import (
          "fmt"
        )

        func Sender(start, quit chan int) {
          x, y := 0, 1
          for { // infinite loop until it break
            select {
            case start <- x:
              x, y = y, x+y //fibnocci
            case <-quit:
              fmt.Println("quit")
              //quit loop when quit chan receives data
              return 		// stop the loop
            default: 		// when no channel receiving data default works
              fmt.Println("no channel is active now")
            }
          }

        }

        func main() {
          var start chan int = make(chan int)
          var quit chan int = make(chan int)
          go func() {
            for i := 0; i < 10; i++ {
              fmt.Println(<-start) // read fib numbers
            }
            quit <- 0 // end after reading 10 numbers
          }()
          Sender(start, quit)
        }


#### using lock mutex

        package main

        import (
          "fmt"
          "sync"
          "time"
        )

        // SafeCounter is safe to use concurrently
        // without lock counter mutiple goroutine manipulate counter
        type SafeCounter struct {
          mu sync.Mutex
          v  map[string]int
        }

        // Inc increments the counter for the given key
        func (c *SafeCounter) Inc(key string) {
          c.mu.Lock()
          // lock so only on gorouting at a time can access map c.v
          c.v[key]++
          c.mu.Unlock()
        }

        // value returns the current value of counter  for given key
        // no worries in this example this func not called by go routines
        // but still we will see how to implement lock
        func (c *SafeCounter) Value(key string) int {
          c.mu.Lock()
          defer c.mu.Unlock() // runs at last
          return c.v[key]
        }

        func main() {
          c := SafeCounter{v: make(map[string]int)}
          for i := 0; i < 1000; i++ {
            go c.Inc("somekey")
          }

          time.Sleep(time.Second) // wait till for loop completes
          fmt.Println(c.Value("somekey"))
        }
        
-------------------


