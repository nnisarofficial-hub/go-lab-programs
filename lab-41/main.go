package main

import "fmt"

func main() {
	arr := [7]int{3, 1, 4, 5, 9, 2, 6}
	min, max := minMax(arr)
	fmt.Println("Numbers:", arr)
	fmt.Println("Minimum: ", min)
	fmt.Println("Maximum: ", max)
}

func minMax(arr [7]int) (min int, max int) {
	min = arr[0]
	max = arr[0]
	for _, val := range arr {
		if val > max {
			max = val
		} else if val < min {
			min = val
		}
	}
	return min, max
}
