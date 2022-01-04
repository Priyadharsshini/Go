/*
CLOSE CHANNELS
we can close a channel so that no more data can be sent through it
*/

package main

import "fmt"

func read(c chan int) {
	fmt.Println(<-c) // reads 5
	fmt.Println(<-c) // blocking call, since read from an empty channel is blocking
}

func main() {
	c := make(chan int)
	go read(c)
	c <- 5   // write operation - Blocking till some other go routine reads from the channel
	close(c) // syntax to close a channel
	c <- 6   // ERROR!!! Cannot put values into a closed channel
	fmt.Println("Program over")

}
