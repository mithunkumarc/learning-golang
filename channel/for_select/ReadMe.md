1. The select statement lets a goroutine wait on multiple communication operations.
2. A select blocks until one of its cases can run, then it executes that case. 
3. It chooses one at random if multiple are ready.
4. use for (infinite/finite) loop to cover multiple cases.
5. use finite for loop for quitting from for-select or dedicated channel for quit from loop.

#### select statement waiting on two channel

    select {
		case rec1 := <-one:
			fmt.Println("rec1 from one channel, msg = ", rec1)
		case rec2 := <-two:
			fmt.Println("rec2 from two channel, msg = ", rec2)
		}

#### select statement with for loop to cover more than one case

    for i := 0; i < 2; i++ {
    		select {
    		case rec1 := <-one:
    			fmt.Println("rec1 from one channel, msg = ", rec1)
    		case rec2 := <-two:
    			fmt.Println("rec2 from two channel, msg = ", rec2)
    		}
    	}

#### quit from for-select using channel

      for {
      		select {
      		case rec1 := <-one:
      			fmt.Println("rec1 from one channel, msg = ", rec1)
      		case rec2 := <-two:
      			fmt.Println("rec2 from two channel, msg = ", rec2)
      		case q := <-quit:
      			fmt.Println("terminating message from channel = ", q)
      			return // quit for loop
      		}
      	}

