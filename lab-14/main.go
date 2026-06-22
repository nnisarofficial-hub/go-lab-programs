package main

import "fmt"

func main() {
	var year int
	fmt.Print("Enter year: ")
	_, err := fmt.Scan(&year)
	if err != nil {
		fmt.Print("Please enter valid year ", err)
		return
	}
	// A year is a leap year if: divisible by 4, AND (not divisible by 100, OR divisible by 400)
	if year%4 == 0 && (year%100 != 0 || year%400 == 0) {
		fmt.Print(year, " is a leap year")
	} else {
		fmt.Print(year, " is not a leap year")
	}
}
