package main

import "fmt"

func main() {
	// addition of 2 numbers
	var (
		input1 int
		input2 int
	)
	fmt.Print("Enter first number: ")
	fmt.Scan(&input1)
	fmt.Print("Enter second number: ")
	fmt.Scan(&input2)

	sum := input1 + input2
	fmt.Printf("The sum is: %d\n", sum)

}
