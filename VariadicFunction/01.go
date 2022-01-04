package main

import "fmt"

//VARIADIC FUNCTION - will take arbitrary number of values.
// the ... behaves like a slice
// Only the last arguement of the function is allowed to be variadic
func add(n ...int) {
	// here ... packs everything into a slice
	sum := 0
	for _, i := range n {
		sum += i
	}
	fmt.Println(sum)
}

func main() {
	add(1)
	add(1, 2, 3, 4, 5)

	n := []int{1, 2, 3, 4, 5, 6, 6, 7, 7555}

	// It will accept array also, and we need to pass the array elements with ...
	// ... is called pack operator.
	// n... means unpack values in n as series of int
	// ... is used both for both pack and unpack
	add(n...)
}

/*
One best example for variadic function is append
func append(slice []type, elems ...Type)[]Type
thats why we give append(slice,1,2,3) or append(slice, anotherslice...)
*/
