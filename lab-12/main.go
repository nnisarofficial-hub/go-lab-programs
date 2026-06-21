package main

import "fmt"

func main() {
	var number float64
	fmt.Print("Enter a number: ")
	_, err := fmt.Scan(&number)
	if err != nil {
		fmt.Println("This is not a number")
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
