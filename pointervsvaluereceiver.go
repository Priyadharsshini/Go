package main

import "fmt"

type person struct {
	name string
	age  int
}

type number []int

func (p *person) changeName(name string) {
	p.name = name
	fmt.Println(p)
}

func (a number) changeNumber(b int) {
	a[2] = b
	fmt.Println(a)
}

func main() {
	p := &person{"priya", 27}
	p.changeName("dharsshini")
	fmt.Println(p)
	t := number{1, 2, 3, 4}
	t.changeNumber(5)
	fmt.Println(t)

}
