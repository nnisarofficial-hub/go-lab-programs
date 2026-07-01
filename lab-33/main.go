package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter text: ")
	text, err := reader.ReadString('\n') // change from '\r' to '\n' so can run on all machines
	if err != nil {
		fmt.Print(err)
		return
	}
	fmt.Print(reverseStringShort(text))
}

func reverseStringShort(s string) string {
	runes := []rune(s)
	slices.Reverse(runes)
	return string(runes)
}
