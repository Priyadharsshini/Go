//PANIC EXAMPLE
package main

import "fmt"

func beforePanic() {
	fmt.Println("I am a deferred function call, and I was scheduled before panicing, so I will be executed in LIFO order")
}

func testRecover() {
	if r := recover(); r != nil {
		fmt.Println("WOHA! Program is panicking with value", r)
	}
	fmt.Println("Done with recovery")
}

func testPanic() {
	fmt.Println("I am going to establish DB connection")
	fmt.Println("DB connection error")
	panic("Not possible to run application")
	fmt.Println("I will not run cos I caused the panic")
}

func main() {
	defer beforePanic()
	defer testRecover()
	testPanic()
	fmt.Println("I will run cos I am recovered")

}
