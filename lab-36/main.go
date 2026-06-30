package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter a sentence: ")
	sentence, err := reader.ReadString('\r')
	if err != nil {
		fmt.Print(err)
		return
	}

	fmt.Printf("Word Count: %d", wordCounter(sentence))
}

func wordCounter(sentence string) int {

	sliceSentence := strings.Fields(sentence)
	return len(sliceSentence)
}
