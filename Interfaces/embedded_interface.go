

package main

import "fmt"

type Shape interface {
	area() float64
}

type Object interface {
	volume() float64
}

type Material interface {
	Shape
	Object
}

type Rect struct {
	length  float64
	breadth float64
}

func (r Rect) area() float64 {
	return r.length * r.breadth
}

func (r Rect) volume() float64 {
	return 87.89
}

func main() {
	var m Material
	r := Rect{65, 78}
	m = r
	fmt.Println(m.area())
}
