package main

import "fmt"

func main() {
	nums := []int{1, 3, 4, 5, 6, 7, 8}

	// Long method
	i := 0
	for i < len(nums) {
		fmt.Print(nums[i], " ")
		i++
	}
	fmt.Println()

	// Short method
	for i := 0; i < len(nums); i++ {
		fmt.Print(nums[i], " ")
	}
	fmt.Println()

	for i := 1; i <= 100; i++ {
		if i%15 == 0 {
			fmt.Println("FizzBuzz")
		} else if i%5 == 0 {
			fmt.Println("Fizz")
		} else if i%3 == 0 {
			fmt.Println("Buzz")
		} else {
			fmt.Println(i)
		}
	}
}
