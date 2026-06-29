package main

import "fmt"

func main() {

	var number int
	fmt.Print("Enter a number: ")
	_, err := fmt.Scan(&number)
	if err != nil {
		fmt.Print(err)
		return
	}
	res := 0
	for number != 0 {
		remainder := number % 10
		res = (res * 10) + remainder
		number /= 10
	}
	fmt.Print("Reversed: ", res)
}
