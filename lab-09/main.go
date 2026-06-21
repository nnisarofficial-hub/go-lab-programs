package main

import "fmt"

func main() {
	var principal, annualRate, time float64
	fmt.Print("Principal (PKR): ")
	fmt.Scan(&principal)

	fmt.Print("Annual rate (%): ")
	fmt.Scan(&annualRate)

	fmt.Print("Time (years): ")
	fmt.Scan(&time)

	simpleInterest := (principal * annualRate * time) / 100
	fmt.Printf("Interest earned: %.2f\nTotal amount: %.2f\n", simpleInterest, simpleInterest+principal)
}
