/*
EMPTY SELECT STATEMENT

Select statement is blocking. when we use empty select, then no other case is there to unblock it, and hence that particular will be blocked resulting in deadlock.

select statement is blocked until one of the cases unblocks, and since there are no case statements available to unblock it, the main goroutine will block forever resulting in a deadlock.
*/

package main

import (
	"fmt"
)

func read(c chan int) {
	fmt.Println(<-c)
}

func main() {
	c := make(chan int)
	go read(c)
	c <- 5
	select {}

}
