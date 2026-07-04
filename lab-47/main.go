package main

import "fmt"

type Rectangle struct {
	Width  float64
	Height float64
}

func main() {
	rectangle1 := Rectangle{Width: 5.0, Height: 3.0}
	fmt.Printf("Rectangle 1: %.1f x %.1f\n", rectangle1.Width, rectangle1.Height)
	fmt.Printf("Area: %.2f\n", rectangle1.Area())
	fmt.Printf("Perimeter: %.2f\n", rectangle1.Perimeter())
	fmt.Printf("Is square: %t\n\n", rectangle1.IsSquare())
	rectangle2 := Rectangle{Width: 4.0, Height: 4.0}
	fmt.Printf("Rectangle 2: %.1f x %.1f\n", rectangle2.Width, rectangle2.Height)
	fmt.Printf("Area: %.2f\n", rectangle2.Area())
	fmt.Printf("Perimeter: %.2f\n", rectangle2.Perimeter())
	fmt.Printf("Is square: %t\n", rectangle2.IsSquare())
}

func (rec Rectangle) Area() float64 {
	return rec.Height * rec.Width
}

func (rec Rectangle) Perimeter() float64 {
	return 2 * (rec.Width + rec.Height)
}

func (rec Rectangle) IsSquare() bool {
	return rec.Height == rec.Width
}
