/* CHANNELS IN GO
Channel is a means of communication between two go routines, where one go routine can send data into channel while another go routine can read data from it
*/

package main

import "fmt"

func read(ch chan int) {
	fmt.Println("M going to read from the channel")
	fmt.Println(<-ch)
}

func main() {
	/*
		Declared with the chan keyword.
		Declaring with 'var' keeps the chan nil.
		Cannot read or write data from nil channel.
	*/

	var c chan int
	fmt.Println(c)

	/*
		we make use of the 'make' to create a ready to use channel.
		Data type of the channel is mandatory, the channel can transport data only of that type.
		Channel is basically a pointer
	*/

	ch := make(chan int)
	fmt.Println(ch)

	/*
		To put some value into a channel
		The below call will be blocked until someother go routine reads from the channel.
		When one goroutine is blocked, go compiler will automatically schedule another go routine.
		This reduces our effort of writing manual locks and releases
	*/

	go read(ch) // At this point we have two go routines.
	fmt.Println("After executing the next line, me(main) go routine will be blocked")
	ch <- 2 // main goroutine is blocked, so the next available go routine will be scheduled
	fmt.Println("The channel was read and I am unblocked")

}
