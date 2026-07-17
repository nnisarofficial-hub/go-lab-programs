package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	menu()
	reader := bufio.NewReader(os.Stdin)
	var todoList []Todo
	for {
		fmt.Print("\n> ")
		userChoice, err := reader.ReadString('\n')
		if err != nil {
			fmt.Print(err)
			return
		}
		userChoice = strings.TrimSpace(userChoice)
		switch userChoice {
		case "1":
			todoList = handleAddTask(reader, todoList)
		case "2":
			listTask(todoList)
		case "3":
			todoList = handleTaskDone(reader, todoList)
		case "4":
			todoList = handleDeleteTask(reader, todoList)
		case "5":
			return
		default:
			fmt.Print("Enter number from 1-5")
		}
	}
}

func menu() {
	fmt.Print("1. Add  2. List  3. Done  4. Delete  5. Quit")
}

type Todo struct {
	ID   int
	Text string
	Done bool
}
// ========================
//
//	Handling Core Functions
//
// ========================
func handleAddTask(reader *bufio.Reader, list []Todo) []Todo {
	fmt.Print("Task: ")
	task, err := reader.ReadString('\n')
	if err != nil {
		fmt.Print(err)
		return list
	}
	task = strings.TrimSpace(task)
	if task == "" {
		fmt.Printf("No task added!")
		return list
	} else {
		list = addTask(list, task)
		fmt.Printf("Added ✓\n")
		return list
	}
}

func handleTaskDone(reader *bufio.Reader, list []Todo) []Todo {
	var id int
	fmt.Print("Mark task ID as done: ")
	_, err := fmt.Scan(&id)
	if err != nil {
		fmt.Print(err)
		return list
	}
	list, found := doneTask(list, id)
	if !found {
		fmt.Printf("Error: Task ID %d not found!\n", id)
	} else {
		fmt.Println("Task marked as done! ✓")
	}
	return list
}

func handleDeleteTask(reader *bufio.Reader, list []Todo) []Todo {
	var id int
	fmt.Print("Enter task ID to delete: ")
	_, err := fmt.Scan(&id)
	if err != nil {
		fmt.Print(err)
		return list
	}
	var found bool
	list, found = deleteTask(list, id)
	if !found {
		fmt.Printf("Error: Task ID %d not found!\n", id)
	} else {
		fmt.Println("Task deleted successfully! ✓")
	}
	return list
}
// ========================
//
//	Core Functions
//
// ========================
func addTask(list []Todo, taskText string) []Todo {
	nextID := len(list) + 1
	newTodo := Todo{ID: nextID, Text: taskText, Done: false}
	return append(list, newTodo)
}

func listTask(list []Todo) {
	if len(list) == 0 {
		fmt.Printf("Task List is Empty")
	}
	for _, task := range list {
		if task.Done {
			fmt.Printf("%d. ✓ %s\n", task.ID, task.Text)
		} else {
			fmt.Printf("%d. ○ %s\n", task.ID, task.Text)
		}
	}
}

func doneTask(list []Todo, targetID int) ([]Todo, bool) {
	for i := range list {
		if list[i].ID == targetID {
			list[i].Done = true
			return list, true
		}
	}
	return list, false
}

func deleteTask(list []Todo, targetID int) ([]Todo, bool) {
	for i := range list {
		if list[i].ID == targetID {
			updatedList := append(list[:i], list[i+1:]...)
			return updatedList, true
		}
	}
	return list, false
}
