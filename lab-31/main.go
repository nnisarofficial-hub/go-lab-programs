package main

import (
	"fmt"
	"math"
)

func main() {
	var number int
	fmt.Print("Enter a number: ")
	_, err := fmt.Scan(&number)
	if err != nil {
		fmt.Print(err)
		return
	}
	if isPrime(number) {
		fmt.Printf("%d is prime ", number)
	} else {
		fmt.Printf("%d is not prime ", number)
	}
}

func isPrime(n int) bool {
	if n <= 1 {
		return false
	}
	if n <= 3 {
		return true
	}
	if n%2 == 0 || n%3 == 0 {
		return false
	}
	limit := int(math.Sqrt(float64(n)))
	for i := 5; i <= limit; i += 6 {
		if n%i == 0 || n%(i+2) == 0 {
			return false
		}
	}
	return true
}
