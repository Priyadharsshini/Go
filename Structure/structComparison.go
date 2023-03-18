package main

import (
	"fmt"
)

type User struct {
	ID       int
	Name     string
	Password string
}

// Two structs are comparable if they belong to the same type and have the same field values.

func main() {
	user1 := User{
		ID:       1,
		Name:     "Priya",
		Password: "Priya",
	}

	user2 := User{
		ID:       1,
		Name:     "Priya",
		Password: "Priya",
	}

	user3 := User{
		ID:       2,
		Name:     "dharsshini",
		Password: "Priya",
	}

	if user1 == user2 {
		fmt.Println("They are equal")
	} else {
		fmt.Println("They are not equal")
	}

	if user2 == user3 {
		fmt.Println("They are equal")
	} else {
		fmt.Println("They are not equal")
	}
}
