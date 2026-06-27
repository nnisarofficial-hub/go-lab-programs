package main

import "fmt"

func main() {
	var counter int
	fmt.Print("Enter N: ")
	_, err := fmt.Scan(&counter)
	if err != nil {
		fmt.Print("Please enter valid number N: ", err)
		return
	}
	for i := 1; i <= counter; i++ {
		fmt.Println(i)
	}
}
