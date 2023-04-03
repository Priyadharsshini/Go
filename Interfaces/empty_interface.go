package main

import "fmt"

type Name struct {
	firstName string
	lastName  string
}

// If an interface is empty, then that means all the type implements that interface.
func explain(i interface{}) {
	fmt.Println("I can accept anything", i)
}

func main() {
	s := "Example"
	t := 5
	n := Name{"John", "Abraham"}
	explain(s)
	explain(t)
	explain(n)

}

/* The above is exactly the way Println works
func Println(a ...interface{})(n int,err error)
- The method has a variadic function and an empty interface which means it accepts any number of any type
- This Println is a perfect example of ploymorphism or you will have to define this same function of all types.
*/
