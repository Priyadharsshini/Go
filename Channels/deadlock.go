/*
DEADLOCK
when one go routine writes to a channel and gets blocked, waiting for some other go routine to read (OR)
when one go routine tries to read an empty channel, then that go routine is blocked

But no other go routines are available to be scheduled, then we enter the state called DEADLOCK.

If you are trying to read data from a channel but channel does not have a value available with it,
it blocks the current goroutine and unblocks other in a hope that some goroutine will push a value to the channel.
Hence, this read operation will be blocking. Similarly, if you are to send data to a channel,
it will block current goroutine and unblock others until some goroutine reads the data from it.
Hence, this send operation will be blocking
*/

package main

func main() {
	c := make(chan int)
	c <- 5

}

//fatal error: all goroutines are asleep - deadlock!
