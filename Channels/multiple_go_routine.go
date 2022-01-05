/*
WORKING WITH MULTIPLE GO ROUTINES
*/
package main

import "fmt"

func square(c chan int) {
	fmt.Println("square starting")
	num := <-c
	c <- num * num
}

func cube(c chan int) {
	fmt.Println("cube starting")
	num := <-c
	c <- num * num * num
}

func main() {
	fmt.Println("M in the main")
	squarechan := make(chan int)
	cubechan := make(chan int)
	go square(squarechan)
	go cube(cubechan)
	squarechan <- 3
	cubechan <- 3
	fmt.Println(<-squarechan + <-cubechan)
}
