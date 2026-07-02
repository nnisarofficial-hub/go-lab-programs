package main

import "fmt"

func main() {
	arrOriginal := []int{1, 2, 3, 4, 5, 1, 2, 5}
	fmt.Println("Original:", arrOriginal)
	fmt.Println("Unique: ", removeDuplicate(arrOriginal))
}

func removeDuplicate(arrOriginal []int) []int {
	arrUnique := []int{}
	seen := make(map[int]bool)
	for _, value := range arrOriginal {
		if !seen[value] {
			seen[value] = true
			arrUnique = append(arrUnique, value)
		}
	}
	return arrUnique
}
