/*
UNIDIRECTIONAL CHANNELS

The use of unidirectional channel is to increase the 'type safety' of a program.
With this we can create a function probably main which can both read and write to a channel, but there is another function which can only read from that channel

We make use of the 'make' function to create the unidirectinal channels

roc := make(<-chan int)
woc := make(chan<- int)

Any write operation on it will result in a fatal error "invalid operation: roc <- "some text" (send to receive-only type <-chan string)".

*/
package main

import "fmt"

// we can also convert bidirectional channels to unidirectional channels like below
func print(c <-chan int) {
	fmt.Println(<-c)
}

func main() {
	fmt.Println("M in the main")
	c := make(chan int)
	go print(c)
	c <- 3
}
