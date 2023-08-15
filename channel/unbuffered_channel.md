1. no memory
2. if you send data to channel and read in the same goroutine , reaches deadlock. because there is no memory to store. to solve this use separate gorouting to send data or use buffered channel.
3. so for unbuffered channel, send data in one goroutine and read in another, it will work.
4. unbuffered channel goes to deadlock in these situation
      1. sending and reading from channel within same goroutine
      2. sending data to channel but not receiving
      3. rading without sending data to channel
5. 
