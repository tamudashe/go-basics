package main

import "fmt"

// Define person struct
type Person struct {
	firstName string
	lastName  string
	city      string
	gender    string
}

// Greeting method: value receiver
func (p Person) greet() string {
	return "Hello, my name is " + p.firstName + " " + p.lastName
}


// Greeting method: pointer receiver
func (p *Person) changeName() {
	p.lastName = "Scott"
}

func main() {
	// Initialize person using struct
	person1 := Person{
		firstName: "Dwight",
		lastName:  "Schrute",
	}

	fmt.Println(person1.firstName, person1.lastName)
	fmt.Println(person1.greet())
	person1.changeName()
	fmt.Println(person1.firstName, person1.lastName)
}
