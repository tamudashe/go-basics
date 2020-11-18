package main

import "fmt"

func main() {
	a := 10
	b := &a

	// use * to read value stored in the memory address
	c := *b
	fmt.Println(a, b, c)

	// change value with pointer
	*b = 20
	fmt.Println(a, b, c)
}
