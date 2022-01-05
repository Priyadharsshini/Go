/*
ADDING TIMEOUT - SELECT STATEMENT

Whenever we have a default case , it is always executed. But we may want that, we need to wait for some time, and within that time, none of the select cases are executed then default case
can be executed. We use the select with timeout for that.

case <-time.After(2 * time.Second)
The above way is the way to add timeout
*/

package main

import (
	"fmt"
	"time"
)

func readvalue(c chan int) {
	c <- 5
}

func read(c chan int) {
	c <- 6
}

func main() {

	c := make(chan int)
	c1 := make(chan int)
	go readvalue(c)
	go read(c1)

	select {
	case num := <-c:
		fmt.Println(num)
	case num1 := <-c1:
		fmt.Println(num1)
	case <-time.After(2 * time.Second):
		fmt.Println("I won let block")
	}

}
