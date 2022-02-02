//PANIC EXAMPLE
package main

import "fmt"

func testPanic() {
	fmt.Println("I am going to establish DB connection")
	fmt.Println("DB connection error")
	panic("Not possible to run application")
}

func main() {
	testPanic()
}
