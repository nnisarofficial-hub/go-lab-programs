package main

import "fmt"

func main() {
	var number int
	fmt.Print("Enter N: ")
	_, err := fmt.Scan(&number)
	if err != nil {
		fmt.Print(err)
		return
	}
	factorial := 1

	if number >= 0 {
		for i := 1; i <= number; i++ {
			factorial *= i
		}
		fmt.Printf("%d! = %d", number, factorial)
	} else {
		fmt.Print("Factorial is not defined for negative numbers")
	}

}
