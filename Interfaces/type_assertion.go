package main

import "fmt"

type Shape interface {
	area() float64
	perimeter() float64
}

type Object interface {
	volume() float64
}

type Rect struct {
	length  float64
	breadth float64
}

func (r Rect) area() float64 {
	return r.length * r.breadth
}

func (r Rect) perimeter() float64 {
	return 2*r.length + 2*r.breadth
}

func (r Rect) volume() float64 {
	return 98
}

func main() {
	var s Shape
	var o Object
	s = Rect{76, 80}
	//This line will give an error "Object does not implement Shape"
	//s = o
	// To do this we need to extract the dynamic value that the interface is pointing to which can be done by type assertion and assign that dynamic type.
	o = s.(Rect)
	fmt.Println(o.volume())
	/* In the above one, if the s does not hold Rect as its dynamic type, then it will give an error as below
	impossible type assertion.
	So before assigning we need to do the right way of type assertion.
	*/
	o, ok := s.(Rect)
	fmt.Println(ok)
	/*In the above syntax, we can check using ok variable if i has concrete type Type or dynamic value of Type.
	If it does not, then ok will be false and value will be the zero value of the Type*/

	// We can convert one interface type to another using type assertion
	new := s.(Object)
	fmt.Println(new)

}
