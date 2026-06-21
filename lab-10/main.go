package main

import "fmt"

func main() {
	var weight, height float64
	fmt.Print("Weight (kg): ")
	fmt.Scan(&weight)

	fmt.Print("Height (m): ")
	fmt.Scan(&height)

	bmi := weight / (height * height)
	if bmi < 18.5 {
		fmt.Printf("BMI: %.2f %s %s\n", bmi, "-", "Underweight")

	} else if bmi >= 18.5 && bmi <= 24.9 {
		fmt.Printf("BMI: %.2f %s %s\n", bmi, "-", "Normal weight")

	} else if bmi >= 25 && bmi <= 29.9 {
		fmt.Printf("BMI: %.2f %s %s\n", bmi, "-", "Overweight")

	} else if bmi >= 30 {
		fmt.Printf("BMI: %.2f %s %s\n", bmi, "-", "Obese")
	}
}
