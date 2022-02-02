package main

import "fmt"

func accessElement(a []int, index int) int {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("You are trying to access an out of bound, so I am going to give the zero value")
		}
	}()
	return a[index]
}

func main() {
	a := []int{1, 2, 3, 4, 5}
	fmt.Println(accessElement(a, 1))
	fmt.Println(accessElement(a, 7))

}
