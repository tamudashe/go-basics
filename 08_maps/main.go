package main

import "fmt"

func main() {
	// Defining a map
	emails := make(map[string]string)
	emails["Dwight"] = "dwight@theoffice.usa"
	emails["Pam"] = "pam@theoffice.usa"
	fmt.Println(emails)

	// Delete from map
	delete(emails, "Pam")
	fmt.Println(emails)

	grades := map[string]int {
		"A": 90,
		"B": 80,
	}
	grades["C"] = 70
	fmt.Println(grades)
}
