package main

import "fmt"

func main() {
	var number int
	fmt.Print("Enter N: ")
	_, err := fmt.Scan(&number)
	if err != nil {
		fmt.Print(err)
		return
	}
	add := 0
	for i := 1; i <= number; i++ {
		add += i
	}
	fmt.Printf("Sum from 1 to %d = %d\n", number, add)
}
