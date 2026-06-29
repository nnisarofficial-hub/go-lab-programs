package main

import (
	"fmt"
	"slices"
)

func main() {

	var word string
	fmt.Print("Enter a word: ")
	if _, err := fmt.Scan(&word); err != nil {
		fmt.Print(err)
		return
	}
	if isPalindrome(word) == true {
		fmt.Print(word, " is a palindrome")
	} else {
		fmt.Print(word, " is not a palindrome")
	}

}

func isPalindrome(word string) bool {
	return word == reverseStringShort(word)
}

func reverseStringShort(sentence string) string {
	runes := []rune(sentence)
	slices.Reverse(runes)
	return string(runes)
}
