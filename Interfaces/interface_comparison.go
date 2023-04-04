/*
INTERFACE COMPARISON
Two interfaces are equal if
- If both the interfaces are null
- If both the interfaces point to the same dynmaic type with the same dynamic value

Even if one interface is null, then they are not equal
*/

package main

import "fmt"

type Shape interface {
	area() float64
}

type Rect struct {
	length  float64
	breadth float64
}

func (r Rect) area() float64 {
	return r.length * r.breadth
}

func main() {
	var s Shape
	s = Rect{10, 30}
	a := s
	if a == s {
		fmt.Println("They are equal")
	} else {
		fmt.Println("No they are not equal")
	}

	var sh Shape
	ah := sh
	if ah == sh {
		fmt.Println("They are equal")
	} else {
		fmt.Println("No they are not equal")
	}

}
