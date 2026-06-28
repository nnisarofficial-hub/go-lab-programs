package main

import "fmt"

func main() {
	var dayNum int
	fmt.Print("Enter day number (1-7): ")
	_, err := fmt.Scan(&dayNum)
	if err != nil {
		fmt.Print("Please enter valid number between (1-7) ", err)
		return
	}
	switch dayNum {
	case 1:
		fmt.Print("Monday")
	case 2:
		fmt.Print("Tuesday")
	case 3:
		fmt.Print("Wednesday")
	case 4:
		fmt.Print("Thursday")
	case 5:
		fmt.Print("Friday")
	case 6:
		fmt.Print("Saturday")
	case 7:
		fmt.Print("Sunday")
	default:
		fmt.Print("Invalid. Please enter a number between 1 and 7.")
	}
}
