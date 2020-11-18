package main

import "fmt"

func main() {
	fmt.Println(greeting("Dwight Schrute"))
	fmt.Println(getSum(44, 22))
}

func greeting(name string) string {
	return "Hello " + name
}

func getSum(num1, num2 int) int {
	return num1 + num2
}