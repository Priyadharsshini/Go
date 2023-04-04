/*
TYPE SWITCH
Take for example an interface is implemented by multiple type and depending on the type that calls the method or the type that is passed as an arguement,
we need to take action accordingly

Consider the example of definining a function with empty interface. In this if the function is called with a string, I want to take a different action, if int then a
different action. In this case we use the type switch

The syntax is i.(type) where i is the interface and the type is a fixed keyword. This will give us the dynamic type of the interface instead of the dynamic value.
*/

package main

import "fmt"

func summary(i interface{}) {
	switch i.(type) {
	case string:
		fmt.Println("I have received a string and I am going to take action accordingly")
	case int:
		fmt.Println("I have received a int and I am going to take action accordingly")
	default:
		fmt.Println("common action for the rest")
	}
}

func main() {
	summary("Apple")
	summary(5)
	summary(34.98)
}
