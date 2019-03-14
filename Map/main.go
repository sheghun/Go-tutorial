package main

import "fmt"

func main() {
	// Syntax #1
	// var colors map[string]string

	// Syntax #2
	// colors := make(map[string]string)

	// Syntax #3
	colors := map[string]string{
		"red":   "#ff0000",
		"green": "#4bf745",
		"white": "#ffffff",
	}

	// Assigning value to a map after initialization
	// colors["white"] = "#ffffff"
	// colors["red"] = "#ff0000"

	// Delete value from a map
	// delete(colors, "red")
	// delete(colors, "white")

	printMap(colors)
}

func printMap(c map[string]string) {
	for color, hex := range c {
		fmt.Println("Hex code for ", color, " is ", hex)
	}
}
