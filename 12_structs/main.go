package main

import "fmt"

// Define person struct
type Person struct {
	firstName string
	lastName  string
	city      string
	gender    string
	age       int32
}

// struct method (value receiver)
func (p Person) greet() string {
	return "hello " + p.firstName
}

func main() {
	// Init Person struct
	person1 := Person{firstName: "Samantha", lastName: "smith", city: "YGN", gender: "m", age: 24}

	fmt.Println(person1)

	fmt.Println(person1.greet())
}
