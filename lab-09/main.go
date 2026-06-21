package main

import "fmt"

func main() {
	var principal float64
	fmt.Print("Principal (PKR): ")
	fmt.Scan(&principal)

	var annualRate float64
	fmt.Print("Annual rate (%): ")
	fmt.Scan(&annualRate)

	var time float64
	fmt.Print("Time (years): ")
	fmt.Scan(&time)

	simpleInterest := (principal * annualRate * time) / 100
	fmt.Printf("Interest earned: %.2f\nTotal amount: %.2f\n", simpleInterest, simpleInterest+principal)
}
