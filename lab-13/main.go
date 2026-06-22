package main

import "fmt"

func main() {
	var num1, num2, num3 float64
	fmt.Print("Enter three numbers: ")
	_, err := fmt.Scan(&num1, &num2, &num3)
	if err != nil {
		fmt.Print("Please enter valid numbers")
		return
	}

	maxVal := num1
	if num2 > maxVal {
		maxVal = num2
	}
	if num3 > maxVal {
		maxVal = num3
	}

	matchCount := 0
	if num1 == maxVal {
		matchCount++
	}
	if num2 == maxVal {
		matchCount++
	}
	if num3 == maxVal {
		matchCount++
	}
	if matchCount > 1 {
		fmt.Print(maxVal, " is the largest (tied)")
	} else {
		fmt.Print(maxVal, " is the largest")
	}
}
