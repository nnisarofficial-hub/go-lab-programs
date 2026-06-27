package main

import "fmt"

func main() {
	var alphabet string
	fmt.Print("Enter a letter: ")
	_, err := fmt.Scan(&alphabet)
	if err != nil {
		fmt.Print("Please enter valid Alphabet: ", err)
	}
	switch alphabet {
	case "A", "a", "E", "e", "I", "i", "O", "o", "U", "u":
		fmt.Print(alphabet, " is a vowel")
	case "B", "b", "C", "c", "D", "d", "F", "f", "G", "g", "H", "h", "J", "j", "K", "k", "L", "l", "M", "m", "N", "n", "P", "p", "Q", "q", "R", "r", "S", "s", "T", "t", "V", "v", "W", "w", "X", "x", "Y", "y", "Z", "z":
		fmt.Print(alphabet, " is a consonant")
	default:
		fmt.Print(alphabet, " Please write an alphabet")
	}
}
