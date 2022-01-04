/*
BUFFERED CHANNEL
Creating a channel without the second parameter in make makes the channel unbuffered.
This means any value written to the channel is immediately available to be read.
This can be changed by creating a channel with a buffer size.
This means that particular channel will be available to be read until after the whole buffer is full.
Read operation on a buffered channel is thirsty, which means read operation will continue until the channel is empty

until the channel received 'n+1' send operation, the channel will not be blocked.
*/

/*
VERY IMPORTANT
goroutine that reads buffer channel will `not block` until the buffer is empty. But it will continue to read
*/

package main

import "fmt"

func read(c chan int) {
	for i := 0; i <= 4; i++ {
		fmt.Println(<-c)
	}
}

func closedwrite(ch chan int) {
	ch <- 1
	ch <- 2
	ch <- 3
	ch <- 4
	ch <- 5
	close(ch)
}

func main() {
	c := make(chan int, 4)
	go read(c)
	c <- 1
	c <- 2
	c <- 3
	c <- 4
	c <- 5 // n+1 operation blocked, 5 will also be read, cos the go routine whichever reads will read until the buffer is empty

	/*
		LENGTH - The number of unread items in the buffered channel
		CAPACITY - The size of the buffer
	*/

	fmt.Println(len(c))
	fmt.Println(cap(c))

	/*
		you can also read values from a buffered channel even before it is full
	*/

	c <- 6
	c <- 7
	fmt.Println(<-c)

	// Buffered channel with close

	ch := make(chan int, 3)
	go closedwrite(ch)

	// even after a channel is closed, data resides in the buffer so we can read it.
	for val := range ch {
		fmt.Println(val)
	}

}
