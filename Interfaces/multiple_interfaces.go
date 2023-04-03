package main

import "fmt"

// A type can implement multiple interfaces.
type Person struct {
	firstName  string
	lastName   string
	age        int
	streeno    string
	streetname string
	province   string
}

type Name interface {
	fullName() string
}

type Address interface {
	fullAddress() string
}

func (p Person) fullName() string {
	return p.firstName + p.lastName
}

func (p Person) fullAddress() string {
	return p.streeno + p.streetname + p.province
}

func main() {
	// In this example we are assigning the struct that has implemented both the interfaces.
	p := Person{"John", "Abraham", 30, "1", "abraham road", "NL"}
	var name Name
	name = p
	var add Address
	add = p
	fmt.Println(name.fullName())
	fmt.Println(add.fullAddress())
	// Now what if we assign one interface to another?. Meaning what will happen if we ask name to hold add
	/*name = add
	fmt.Println(name.fullName())
	The above will throw an error cos the address does not implement the Name. They both are static types and they are different and only the underlying type they
	both point to are the same.
	You will get the following error
	./prog.go:41:9: cannot use add (variable of type Address) as Name value in assignment: Address does not implement Name (missing method fullName)
	*/

}
