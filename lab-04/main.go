package main

import "fmt"

func main() {
	var num1, num2 float64
	var operator string

	fmt.Print("Enter first number: ")
	fmt.Scan(&num1)

	fmt.Print("Enter operator (+, -, *, /): ")
	fmt.Scan(&operator)

	fmt.Print("Enter second number: ")
	fmt.Scan(&num2)

	switch operator {
	case "+":
		fmt.Print(num1, " + ", num2, " = ", num1+num2)
	case "-":
		fmt.Print(num1, " - ", num2, " = ", num1-num2)
	case "*":
		fmt.Print(num1, " * ", num2, " = ", num1*num2)
	case "/":
		if num2 != 0 {
			fmt.Print(num1, " / ", num2, " = ", num1/num2)
		} else {
			fmt.Print("Error: cannot divide by zero")
		}
	default:
		fmt.Print("Unknown operator. Please use operator (+, -, *, /): ")
	}

}
