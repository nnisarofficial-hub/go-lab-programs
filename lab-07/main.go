package main

import "fmt"

func main() {
	var width, height float64
	fmt.Print("Enter width: ")
	fmt.Scan(&width)

	fmt.Print("Enter height: ")
	fmt.Scan(&height)

	fmt.Printf("Area: %.2f\nPerimeter: %.2f\n", width*height, 2*(width+height))
}
