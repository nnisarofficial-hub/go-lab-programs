package main

import "fmt"

// I take help from the internet but I really understood by myself i can write code again without taking any help
// and please remove these comments while checking

func main() {
	var star int
	fmt.Print("Enter number: ")
	fmt.Scan(&star)
	for i := 1; i <= star; i++ {
		for j := 1; j <= star-i; j++ {
			fmt.Print(" ")
		}
		for k := 1; k <= i; k++ {
			fmt.Print("*")
		}
		fmt.Println()
	}

}
