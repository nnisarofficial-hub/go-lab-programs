package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strings"
)

func main() {

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter text: ")
	text, err := reader.ReadString('\n')
	if err != nil {
		fmt.Print(err)
		return
	}
	cleanedText := strings.TrimSpace(text)
	fmt.Print(reverseString(cleanedText))
}

func reverseString(s string) string {
	runes := []rune(s)
	slices.Reverse(runes)
	return string(runes)
}
