/*
CHANNELS - SELECT STATEMENT

DEFAULT CASE
Default make the select case always non-blocking.
If any of the select case has a non-blocking or value readily available, then that will be executed.If not default case will be executed.
Won have any chance to block that routine and schedule another routine.

Default cases are very useful when we want to avoid deadlock, since it is non-blocking.

*/



package main

import "fmt"

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
	default:
		fmt.Println("I won let block")
	}

}
