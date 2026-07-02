package main

import "fmt"

func main() {
	arr := []float64{85.0, 92.5, 78.0, 90.0, 88.5}
	sum, count, average := sliceAvg(arr)
	fmt.Println("Numbers:", arr)
	fmt.Printf("Sum: %.2f\n", sum)
	fmt.Println("Count:", count)
	fmt.Printf("Average: %.2f\n", average)
}
func sliceAvg(arr []float64) (sum float64, count int, average float64) {
	for _, val := range arr {
		sum += val
		count += 1
	}
	average = sum / float64(count)
	return sum, count, average
}
