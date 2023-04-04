/*
TYPE ASSERTION
i.(Type) -> this is how we use type assertion, where 'i' is the interface and the 'Type' is the dynamic type you think the interface implements
and whose dynamic value you want to extract
*/

package main

import "fmt"

type Name interface {
	fullName() string
}

type Address interface {
	fullAddress() string
}

type Person struct {
	firstName string
	lastName  string
}

type Location struct {
	streetName string
	streetNo   string
	province   string
}

func (p Person) fullName() string {
	return p.firstName + p.lastName
}

func main() {
	var n Name = Person{"John", "Abraham"}
	// Here we are using type assertion to extract the dynamic value that our interface is pointing to
	s := n.(Person)
	fmt.Println(s.fullName())
	/*
		- Suppose our interface does not point to the dynamic type at all(not implemented)
		- or Suppose it pointing to the dynamic type, but the values are null, then it will run into run time error
		To overcome this we have the below syntax, to fail silently
	*/
	a, ok := n.(Person)
	if ok {
		fmt.Println(a)
		fmt.Println("Suceeded")
	}
	// We can also use another interface in the 'Type' place, in that case it will check if the 'i' dynamic type implements the 'Type' too
	b, ok := n.(Address)
	if !ok {
		fmt.Println(b)
		fmt.Println("The underslying dynamic type in 'n' does not implement Address")
	}
	// In the same above way we can convert one interface type to another.

}
