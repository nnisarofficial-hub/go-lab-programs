package main

import "fmt"

func main() {
	input := ""
	fmt.Print("Enter one word: ")

	fmt.Scan(&input)

	fmt.Print("Hello, ", input, "!")
}
