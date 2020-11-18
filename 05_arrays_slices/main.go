package main

import "fmt"

func main() {
	// Arrays: fixed length and types have to be named
	var fruitArr [2]string

	// Assign values to the array
	fruitArr[0] = "apple"
	fruitArr[1] = "orange"
	fmt.Printf("%T\n", fruitArr)

	// Declare and assign values to the array
	numbers := [2]int{1, 4}
	fmt.Println(numbers)

	// Slices
	grades := []int{98, 100, 29, 40}
	fmt.Printf("%T\n", grades)
}
