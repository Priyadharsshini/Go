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

func main() {
	c := make(chan int)
	go generateCube(c)
	for {
		if val, ok := <-c; ok {
			fmt.Println(val)
		} else {
			break
		}
	}

}
