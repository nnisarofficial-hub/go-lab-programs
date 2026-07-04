package main

import "fmt"

func main() {
	studentData()
}

type student struct {
	Name  string
	Score float64
}

func studentData() {
	studentsData := []student{
		{Name: "Ali", Score: 88.5},
		{Name: "Sara", Score: 95.0},
		{Name: "Umar", Score: 72.0},
		{Name: "Fatima", Score: 61.5},
		{Name: "Hassan", Score: 45.0},
	}
	finalGrades := grades(studentsData)
	var totalScore float64
	fmt.Printf("%-12s %-8s %s\n", "Name", "Score", "Grade")
	for i, std := range studentsData {
		scoreStr := fmt.Sprintf("%.1f", std.Score)
		fmt.Printf("%-12s %-8s %s\n", std.Name, scoreStr, finalGrades[i])
		totalScore += std.Score
	}
	fmt.Printf("\nClass average: %.2f\n", totalScore/float64(len(studentsData)))
}

func grades(list []student) []string {
	var gradesList []string
	for _, std := range list {
		var letterGrade string
		if std.Score >= 90 && std.Score <= 100 {
			letterGrade = "A"
		} else if std.Score >= 80 && std.Score <= 89 {
			letterGrade = "B"
		} else if std.Score >= 70 && std.Score <= 79 {
			letterGrade = "C"
		} else if std.Score >= 60 && std.Score <= 69 {
			letterGrade = "D"
		} else if std.Score < 60 && std.Score >= 0 {
			letterGrade = "F"
		} else if std.Score > 100 || std.Score < 0 {
			fmt.Print("Invalid score. Must be between 0 and 100.")
		}
		gradesList = append(gradesList, letterGrade)
	}
	return gradesList
}
