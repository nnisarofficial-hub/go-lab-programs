package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter sentence: ")
	sentence, err := reader.ReadString('\n')
	if err != nil {
		fmt.Print(err)
		return
	}
	result := wordFrequency(sentence)
	for key, value := range result {
		fmt.Printf("%s: %d\n", key, value)
	}
}

func wordFrequency(sentence string) map[string]int {
	lowerCaseSentence := strings.ToLower(sentence)
	words := strings.Fields((lowerCaseSentence))
	wordCounts := make(map[string]int)
	for _, word := range words {
		cleanedWord := strings.TrimFunc(word, func(r rune) bool {
			return unicode.IsPunct(r)
		})
		if cleanedWord != "" {
			wordCounts[cleanedWord]++
		}
	}
	return wordCounts
}
