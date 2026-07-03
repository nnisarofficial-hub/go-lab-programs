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
	sentence, err := reader.ReadString('\n')
	if err != nil {
		fmt.Print(err)
		return
	}
	sentence = strings.TrimSpace(sentence)
	fmt.Print("Enter Character: ")
	character, err := reader.ReadString('\n')
	if err != nil {
		fmt.Print(err)
		return
	}
	character = strings.TrimSpace(character)
	fmt.Printf("'%s' appears %d times in \"%s\"", character, countChar(sentence, character), sentence)
}

func countChar(sentence, character string) int {
	return strings.Count(sentence, character)
}
