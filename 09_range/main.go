package main

import (
	"fmt"
)

func main() {
	nums := []int{33, 45, 3, 67, 10, 7, 8, 5}
	result := 0
	for _, num := range nums {
		result += num
	}
	fmt.Println(result)

	grades := map[string]int{
		"A": 90,
		"B": 80,
		"C": 70,
		"D": 60,
		"F": 50,
	}

	for grade, percent := range grades {
		fmt.Printf("%s: %d\n", grade, percent)
	}
	for key := range grades {
		fmt.Println(key)
	}
}
