/*
ANONYMOUS GO ROUTINE
*/

package main

import "fmt"

func main() {
	c := make(chan int)
	go func(c chan int) {
		fmt.Println(<-c)
	}(c)
	c <- 3
	fmt.Println("main over")
}
