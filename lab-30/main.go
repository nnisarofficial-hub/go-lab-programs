package main

import "fmt"

func main() {
	const number = 42
	counter := 0
	for {
		var guess int
		fmt.Print("Guess the number: ")
		_, err := fmt.Scan(&guess)
		if err != nil {
			fmt.Print(err)
			return
		}
		if guess > number {
			fmt.Println("Too high")
			counter += 1
		} else if guess < number {
			fmt.Println("Too low")
			counter += 1
		} else if guess == number {
			counter += 1
			fmt.Printf("Correct! You got it in %d attempts.", counter)
			break
		}
	}
}
