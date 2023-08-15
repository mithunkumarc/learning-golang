1. no memory
2. if you send data to channel and read in the same goroutine , reaches deadlock. because there is no memory to store.
3. so for unbuffered channel, send data in one goroutine and read in another, it will work.
