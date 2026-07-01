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
	fmt.Print(removeSpaces(sentence))
	fmt.Printf("Removed %d spaces", countSpaces(sentence))

}

func countSpaces(sentence string) int {
	return strings.Count(sentence, " ")
}
func removeSpaces(sentence string) string {
	result := strings.ReplaceAll(sentence, " ", "")
	return result
}
