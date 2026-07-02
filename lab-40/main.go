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
<<<<<<< HEAD
	sentence = strings.TrimSpace(sentence)
	fmt.Print("Enter Character: ")
	character, err1 := reader.ReadString('\n')
	if err1 != nil {
		fmt.Print(err1)
		return
	}
	character = strings.TrimSpace(character)
	fmt.Printf("'%s' appears %d times in %s", character, countChar(sentence, character), sentence)
}
=======
	fmt.Print("Enter Character: ")
	character, err := reader.ReadString('\n')
	if err != nil {
		fmt.Print(err)
		return
	}
	character = strings.TrimSpace(character)
	fmt.Printf("'%s' appears %d times in \"%s\"", character, countChr(sentence, character), sentence)
}

>>>>>>> 862bb0b5c10b78c2bcbd741177f88df200ce0208
func countChar(sentence, character string) int {
	return strings.Count(sentence, character)
}
