package main

import "fmt"

func main() {
	var number int
	fmt.Print("Enter a number: ")
	_, err := fmt.Scan(&number)
	if err != nil {
		fmt.Print(err)
	}

	res := 0
	for number != 0 {
		remainder := number % 10
		res += remainder
		number /= 10
	}
	fmt.Print(res)
}
