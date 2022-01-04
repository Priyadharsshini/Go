/*
CLOSE CHANNELS
we can close a channel so that no more data can be sent through it
*/

package main

import "fmt"

func generateCube(c chan int) {
	for i := 0; i <= 10; i++ {
		c <- i * i * i
	}
	close(c) // If I don't close here then deadlock
}

/*
In the previous program we manually checked for the closing of the channel.
To avoid that go lang provides a much easier option.
when you use range, it will automatically check for closed channel and break automatically.
*/

func main() {
	c := make(chan int)
	go generateCube(c)
	for val := range c {
		fmt.Println(val)
	}

}
