package main

import "fmt"

func main() {
	var baseValue int
	fmt.Print("Enter base: ")
	if _, err := fmt.Scan(&baseValue); err != nil {
		fmt.Print(err)
		return
	}
	var expValue int
	fmt.Print("Enter exponent: ")
	if _, err := fmt.Scan(&expValue); err != nil {
		fmt.Print(err)
		return
	}
	fmt.Printf("%d^%d = %d", baseValue, expValue, power(baseValue, expValue))
}

func power(baseValue, expValue int) int {
	valueStore := 1
	for i := 1; i <= expValue; i++ {
		valueStore *= baseValue
	}
	return valueStore
}
