/*
NIL CHANNEL - SELECT STATEMENT

no values can be read or write to a nil channel. So it will gives nil channel error
'goroutine 1 [chan receive (nil chan)]'


secondly nil channel on a select is considered as an empty select. Deadlock
'fatal error: all goroutines are asleep - deadlock!'
*/

package main

import (
	"fmt"
)

func write(c chan int) {
	c <- 5
}

func main() {
	var c chan int
	go write(c)
	select {
	case num := <-c:
		fmt.Println(num)
	}

}
