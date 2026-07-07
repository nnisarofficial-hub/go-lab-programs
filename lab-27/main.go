package main

import (
	"fmt"
	"strings"
)

func main() {
	var rows int
	fmt.Print("Enter the number of rows: ")
	fmt.Scanln(&rows)
	for i := 1; i <= rows; i++ {
		spaces := strings.Repeat(" ", rows-i)
		stars := strings.Repeat("*", i)
		fmt.Println(spaces + stars)
	}
}
