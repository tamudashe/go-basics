package main

import "fmt"

func main() {

	color := "redd"
	// Switch
	switch color {
	case "red":
		fmt.Println("the color is red")
	default:
		fmt.Println("invalid color")
	}
}
