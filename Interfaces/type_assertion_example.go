type Animal interface {
    Speak() string
}

type Dog struct {
    Name string
}

func (d Dog) Speak() string {
    return "Woof!"
}

type Cat struct {
    Name string
}

func (c Cat) Speak() string {
    return "Meow!"
}

func main() {
    animals := []Animal{Dog{Name: "Fido"}, Cat{Name: "Whiskers"}}

    for _, animal := range animals {
        sound := animal.Speak()
        fmt.Println(sound)

        // type assertion - Based on the dynamic type, provide more information
        if dog, ok := animal.(Dog); ok {
            fmt.Println(dog.Name, "is a dog")
        } else if cat, ok := animal.(Cat); ok {
            fmt.Println(cat.Name, "is a cat")
        }
    }
}
