package main

import "fmt"

type Shape interface {
	Area() float64
}

type Circle struct {
	radius float64
}

type Square struct {
	side float64
}

type Rectangle struct {
	length  float64
	breadth float64
}

func (c Circle) Area() float64 {
	return 3.14 * c.radius * c.radius
}

func (s Square) Area() float64 {
	return s.side * s.side
}

func (r Rectangle) Area() float64 {
	return r.length * r.breadth
}

func main() {
	c := Circle{10}
	s := Square{5}
	r := Rectangle{2, 5}
	shapes := []Shape{c, s, r}
	for _, shape := range shapes {
		fmt.Println(shape.Area())
	}
}
