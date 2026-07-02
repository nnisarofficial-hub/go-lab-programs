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
	sentence = strings.TrimSpace(sentence)
	fmt.Print("Enter Character: ")
	character, err1 := reader.ReadString('\n')
	if err1 != nil {
		fmt.Print(err1)
		return
	}
	character = strings.TrimSpace(character)
	fmt.Printf("'%s' appears %d times in %s", character, countChr(sentence, character), sentence)
}
func countChr(sentence, character string) int {
	return strings.Count(sentence, character)
}
