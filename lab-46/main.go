package main

import "fmt"

func main() {
	studentsData := []Student{
		{Name: "Ali", Score: 88.5},
		{Name: "Sara", Score: 95.0},
		{Name: "Umar", Score: 72.0},
		{Name: "Fatima", Score: 61.5},
		{Name: "Hassan", Score: 45.0},
	}
	var totalScore float64
	fmt.Printf("%-12s %-8s %s\n", "Name", "Score", "Grade")
	for _, std := range studentsData {
		fmt.Printf("%-12s %-8.1f %s\n", std.Name, std.Score, std.Grade())
		totalScore += std.Score
	}
	fmt.Printf("\nClass average: %.2f\n", totalScore/float64(len(studentsData)))
}

type Student struct {
	Name  string
	Score float64
}

func (s Student) Grade() string {
	if s.Score >= 90 && s.Score <= 100 {
		return "A"
	} else if s.Score >= 80 && s.Score <= 89 {
		return "B"
	} else if s.Score >= 70 && s.Score <= 79 {
		return "C"
	} else if s.Score >= 60 && s.Score <= 69 {
		return "D"
	} else if s.Score < 60 && s.Score >= 0 {
		return "F"
	} else if s.Score > 100 || s.Score < 0 {
		return "Invalid score. Must be between 0 and 100."
	}
	return "Invalid"
}
