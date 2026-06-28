package main

import "fmt"

func main() {
	var fibNum int
	fmt.Print("How many Fibonacci numbers? ")
	_, err := fmt.Scan(&fibNum)
	if err != nil {
		fmt.Print(err)
		return
	}
	prev, current := 0, 1
	for i := 1; i <= fibNum; i++ {
		fmt.Printf("%d ", prev)
		prev, current = current, prev+current

	}
}
