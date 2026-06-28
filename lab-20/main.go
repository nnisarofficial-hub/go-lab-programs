package main

import "fmt"

func main() {
	var length1, length2, length3 int
	fmt.Print("Enter three sides: ")
	_, err := fmt.Scan(&length1, &length2, &length3)
	if err != nil {
		fmt.Print("Please enter valid lengths: ", err)
		return
	}

	if length1 >= (length2+length3) || length2 >= (length1+length3) || length3 >= (length1+length2) {
		fmt.Print("Not a valid triangle")
	} else if length1 == length2 && length2 == length3 {

		fmt.Print("Equilateral triangle")
	} else if length1 == length2 || length2 == length3 || length1 == length3 {
		fmt.Print("Isosceles triangle")
	}
}
