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
	if _, err1 := fmt.Scan(&expValue); err1 != nil {
		fmt.Print(err1)
		return
	}
	fmt.Print(power(baseValue, expValue))
}

func power(baseValue, expValue int) int {

	valueStore := 1
	for i := 1; i <= expValue; i++ {
		valueStore *= baseValue
	}
	return valueStore

}
