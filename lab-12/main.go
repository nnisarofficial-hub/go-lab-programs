package main

import "fmt"

func main() {
	var number float64
	fmt.Print("Enter a number: ")
	_, err := fmt.Scan(&number)
	if err != nil {
		fmt.Printf("This is not a number: %v", err)
		return
	}

	if number > 0 {
		fmt.Print(number, "is positive")
	} else if number < 0 {
		fmt.Print(number, " is negative")
	} else {
		fmt.Print(number, " is zero")
	}

}
