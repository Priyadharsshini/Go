/*
CHANNELS - SELECT STATEMENT

Select statement is like a switch case except that it is used in channels only.
Select operation is blocking. (except when we have a DEFAULT case)
When all the select cases are blocking by default, then select will block for sometime, till any of the case becomes unblocked and that particular case will be exceuted.
If one or more case is a non-blocking case, then select will execute any one of the non blocking randomly.

*/

/* 
USE CASE 
Above program simulates real world web service where a load balancer gets millions of requests and it has to return a response from one of the services available. 
Using goroutines, channels and select, we can ask multiple services for a response, and one which responds quickly can be used.
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
	}

}
