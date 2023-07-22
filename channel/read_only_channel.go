/*
A channel can be made read-only to whoever receives it, while the sender still has a two-way channel to which they can write. 
For example:
*/

func F() <-chan int {
    // Create a regular, two-way channel.
    c := make(chan int)

    go func() {
        defer close(c)

        // Do stuff
        c <- 123
    }()

    // Returning it, implicitly converts it to read-only,
    // as per the function return type.
    return c
}
/*
Whoever calls F(), receives a channel from which they can only read. 
This is mostly useful to avoid potential misuse of a channel at compile time. 
Because read/write-only channels are distinct types, the compiler can use its existing type-checking mechanisms to ensure the 
caller does not try to write stuff into a channel it has no business writing to.
https://stackoverflow.com/questions/13596186/whats-the-point-of-one-way-channels-in-go
*/
