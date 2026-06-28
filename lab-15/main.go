package main

import "fmt"

func main() {
	var score float64
	fmt.Print("Enter score: ")
	_, err := fmt.Scan(&score)
	if err != nil {
		fmt.Print("Please enter valid score ", err)
		return
	}

	if score >= 90 && score <= 100 {
		fmt.Print("Grade: A")
	} else if score >= 80 && score <= 89 {
		fmt.Print("Grade: B")
	} else if score >= 70 && score <= 79 {
		fmt.Print("Grade: C")
	} else if score >= 60 && score <= 69 {
		fmt.Print("Grade: D")
	} else if score < 60 && score > 0 {
		fmt.Print("Grade: F")
	} else if score > 100 || score < 0 {
		fmt.Print("Invalid score. Must be between 0 and 100.")
	}

}
