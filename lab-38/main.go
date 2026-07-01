package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter message: ")
	message, err := reader.ReadString('\n')
	if err != nil {
		fmt.Print(err)
		return
	}
	var shift int
	fmt.Print("Enter shift: ")
	_, err1 := fmt.Scanf("%d\n", &shift)
	if err1 != nil {
		fmt.Print(err1)
		return
	}
	fmt.Print(caesarCipher(message, shift))
}
func caesarCipher(message string, shift int) string {

	encryption := make([]rune, 0)

	for _, i := range message {
		if i >= 'A' && i <= 'Z' {
			encryptedUpper := 'A' + (i-'A'+rune(shift))%26
			encryption = append(encryption, encryptedUpper)
		} else if i >= 'a' && i <= 'z' {
			encryptedLower := 'a' + (i-'a'+rune(shift))%26
			encryption = append(encryption, encryptedLower)
		} else {
			encryption = append(encryption, i)
		}
	}
	return string(encryption)
}
