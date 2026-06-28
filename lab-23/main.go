package main

import "fmt"

func main() {
	var number int
	fmt.Print("Enter number: ")
	_, err := fmt.Scan(&number)
	if err != nil {
		fmt.Print(err)
		return
	}
	for i := 1; i <= 10; i++ {
		fmt.Printf("%d x %2d = %3d\n", number, i, number*i)
	}
}
