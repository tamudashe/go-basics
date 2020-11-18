package main

import "fmt"

func main() {
	// Main types
	// string
	// bool
	// int
	// int int8 int16 int32 int64
	// uint uint8 uint16 uint64 uintptr
	// byte - alias for uint8
	// rune - alias for int32
	// float32 float64
	// complex64 complex128 -> really large numberes

	// create a variable using the var keyword
	var name = "Tamudashe"
	const year = 2020
	const isCool = false
	fmt.Println(name, year)
	fmt.Printf("%T\n", isCool)

	// shorthand method, can only be used within a function
	quote := "to be or not to be"
	fmt.Println(quote)

	size := 1.3
	fmt.Printf("%T\n", size)

	name, email := "Dwight", "dwight@theofficeusa.com"
	fmt.Println(name, email)

}
