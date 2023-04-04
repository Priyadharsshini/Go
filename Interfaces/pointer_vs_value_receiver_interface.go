/*
POINTER VS VALUE RECEIVER IN INTERFACES
*/

package main

import "fmt"

type Shape interface {
	area() float64
	perimeter() float64
}

type Rect struct {
	length  float64
	breadth float64
}

func (r *Rect) area() float64 {
	return r.length * r.breadth
}

func (r Rect) perimeter() float64 {
	return 2*r.length + 2*r.breadth
}

func main() {
	r := Rect{10.98, 6}
	/* The Below will fail saying the following
	./prog.go:29:16: cannot use r (variable of type Rect) as Shape value in variable declaration: Rect does not implement Shape (method area has pointer receiver)
	*/
	//var s Shape = r
	//unlike struct a menthod with pointer receiver will work both on pointer and value, cause automaric type conversions happen. but here it does not work like that.
	var s Shape = &r // we need to pass the address, pointer.
	fmt.Println(s.area())

	//However there is another catch
	fmt.Println(s.perimeter())
	// The above will work even though we are passing address but the perimeter accepts only value receiver , but looks like go compiler automatically passes a copy to that.
}
