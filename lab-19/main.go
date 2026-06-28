package main

import (
	"fmt"
	"strings"
)

func main() {
	var alphabet string
	fmt.Print("Enter a letter: ")
	fmt.Scan(&alphabet)

	lowerCase := strings.ToLower(alphabet)
	switch lowerCase {
	case "a", "e", "i", "o", "u":
		fmt.Print(alphabet, " is a vowel")
	case "b", "c", "d", "f", "g", "h", "j", "k", "l", "m", "n", "p", "q", "r", "s", "t", "v", "w", "x", "y", "z":
		fmt.Print(alphabet, " is a consonant")
	default:
		fmt.Print(alphabet, " Please write an alphabet")
	}
}
