package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter sentence: ")
	sentence, err := reader.ReadString('\r')
	if err != nil {
		fmt.Print(err)
		return
	}

	fmt.Printf("Vowels: %d", countVowels(sentence))
}
func countVowels(senence string) int {

	lowerStr := strings.ToLower(senence)
	counter := 0
	for _, char := range lowerStr {
		if char == 'a' || char == 'e' || char == 'i' || char == 'o' || char == 'u' {
			counter++
		}
	}
	return counter
}
