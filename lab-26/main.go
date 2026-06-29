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
	for i := number; i > 0; i-- {
		fmt.Println(i)
	}
	fmt.Println("Go!")
}
