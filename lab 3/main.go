package main

import "fmt"

func main() {
	// addition of 2 numbers
	var num1, num2 int
	fmt.Print("Enter first number: ")
	fmt.Scan(&num1)
	fmt.Print("Enter second number: ")
	fmt.Scan(&num2)

	fmt.Printf("%d + %d = %d\n", num1, num2, num1+num2)

}
