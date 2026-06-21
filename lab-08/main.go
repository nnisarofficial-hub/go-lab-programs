package main

import (
	"fmt"
	"math"
)

func main() {
	var radius float64
	fmt.Print("Enter radius: ")
	fmt.Scan(&radius)

	fmt.Printf("Area: %.2f\nCircumference: %.2f\n", math.Pi*math.Pow(radius, 2), 2*math.Pi*radius)
}
