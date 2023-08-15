#### document : https://www.youtube.com/watch?v=VkGQFFl66X4

#### unbuffered channel 

	cannot execute in the same goroutine? https://www.youtube.com/watch?v=VkGQFFl66X4

#### Unbuffered Channel(default capacity 1)
	
     it is synchronous, when data is pushed to unbuffered channel, 
     it will wait until this data is read(blocking)
     when the data is read, channel opens up again for next data 

#### Buffered Channel(while create we mention the capacity of buffer)
	
     Example: Assume capacity of channel is 3
  	it is asynchronous until the capacity is full
  	when the first data is pushed to Buffered Channel, since channel has capacity of 3, 
   	still 2 slot left and channel is not blocked and ready for next new data.
  	when channel capacity is full, and no one read the data yet from this channel then channel gets blocked.
  	when data is read from the channel then channel opens up again for data as capacity is 3-1 this time. 
   	vice versa
  	
  	https://www.youtube.com/watch?v=qyM8Pi1KiiM

#### what is readonly channel?

#### read

	https://go101.org/article/channel.html
 	https://golangdocs.com/channels-in-golang
 	https://pkg.go.dev/github.com/eapache/channel
