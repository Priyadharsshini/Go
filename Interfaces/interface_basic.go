/*
Interface -> implement polymorphism:

Code reusability: you can define a set of method signatures that multiple types can implement. 
This allows you to write code that can work with different types that share the same interface, without having to rewrite the code for each individual type.

Abstraction: you can abstract away the implementation details of the types that implement the interface. 
This allows you to write code that is more generic and can be used with different types, without needing to know the specific implementation details of each type.

Flexibility: Polymorphism allows you to write code that is more flexible and can work with different types that share the same interface. 
This means that you can add new types that implement the interface and use them interchangeably with the existing types, 
without having to change the code that uses them.

Testability: Interfaces make it easier to test your code, as you can create mock implementations of the interface to test your code with, 
rather than having to rely on real implementations of the types that implement the interface.
*/

/*
1) Collection of method signatures that a Type can implement.
2) It just provides the method name, input arguements and the return type(if any)
3) If a type implements the method signature provided by the interface, then the type is said to implement the interface.
*/

// NOTE: You must implement all the method declared in the interface for a type to implement the interface.
package main

import "fmt"

type Shape interface {
	area() float64
	perimeter() float64
}

// We are creating a Rect type that implements the Shape type.
type Rect struct {
	length float64
	width  float64
}

type Square struct {
	side float64
}

func (s Square) area() float64 {
	return s.side * s.side
}

func (s Square) perimeter() float64 {
	return 4 * s.side
}

func (r Rect) area() float64 {
	return r.length * r.width
}

func (r Rect) perimeter() float64 {
	return 2*r.length + 2*r.width
}

func main() {
	// Since Interface is just a type we can create variables of the type.
	var s Shape
	/*
		Interface usually have two types
		- Static type - This is the static type of the interface. Here in this example, it is the shape
		- Dynamic type - Usually, interface always holds a dynamic value and the type of the interface becomes the type that the dynamic value belongs
	*/
	s = Rect{20, 30}
	fmt.Println(s)
	fmt.Println(s.area())
	fmt.Println(s.perimeter())
	// Now the dynamic value and type of Shape interface has been changed to Square
	s = Square{10}
	fmt.Println(s)
	fmt.Println(s.area())
	fmt.Println(s.perimeter())
	// We can also call rectange with rectange type directly, like how we normally do and similary for square too
	re := Rect{40, 70}
	fmt.Println(re.area())
	fmt.Println(re.perimeter())
	//Interface works in a ways by dynamically holding reference to an underlying type
}
