package main

import "fmt"

// The first main use of interface is for abstraction. To hide the implmentation detail. 
// The second use is to avoid code deuplication and to facilitate easy extensibility.

//Let us assume we have a fly function for birds to fly

type Bird struct{
	name String
}

func fly(b Bird){
	fmt.Println("I am flying", name)
}

// Suppose next comes where you want to have helicopter that can fly, you need to change the above code function name to flyBird -> modify existing code.

func main(){
	b := Bird
	fly(b)
	// for helicoper again, then any other type same. 
	// But with interface 
	//b := Bird
	//h := helicopter
	//f := flight
	//var common Common
	//common = b //we will be changing only this in the main
	//common.fly()
