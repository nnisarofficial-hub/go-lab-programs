package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Todo struct {
	ID   int
	Text string
	Done bool
}

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
		choice := strings.TrimSpace(userChoice)
		switch choice {
		case "1":
			todoList = handleAddTask(reader, todoList)
		case "2":
			handleListTask(todoList)
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
	taskClean := strings.TrimSpace(task)
	if taskClean == "" {
		fmt.Printf("No task added!")
		return list
	} else {
		list = addTask(list, taskClean)
		fmt.Printf("Added ✓\n")
		return list
	}
}

func handleListTask(list []Todo) {
	if len(list) == 0 {
		fmt.Println("No Task Found")
		return
	}
	listTask(list)
}

func handleTaskDone(reader *bufio.Reader, list []Todo) []Todo {
	fmt.Print("Mark task ID as done: ")
	id, err := reader.ReadString('\n')
	if err != nil {
		fmt.Print(err)
		return list
	}
	idClean := strings.TrimSpace(id)
	targetId, err := strconv.Atoi(idClean)
	if err != nil {
		fmt.Print("Please enter valid number ID! ", err)
		return list
	}
	list = doneTask(list, targetId)
	return list
}

func handleDeleteTask(reader *bufio.Reader, list []Todo) []Todo {
	fmt.Print("Enter task ID to delete: ")
	id, err := reader.ReadString('\n')
	if err != nil {
		fmt.Print(err)
		return list
	}
	idClean := strings.TrimSpace(id)
	targetId, err := strconv.Atoi(idClean)
	if err != nil {
		fmt.Print("Please enter valid number ID! ", err)
		return list
	}
	list = deleteTask(list, targetId)
	return list
}

// ========================
//
//	Core Functions
//
// ========================
func addTask(list []Todo, taskTest string) []Todo {
	nextID := len(list) + 1
	newTodo := Todo{ID: nextID, Text: taskTest, Done: false}
	return append(list, newTodo)
}

func listTask(list []Todo) {
	for _, job := range list {
		if job.Done {
			fmt.Printf("%d. ✓ %s\n", job.ID, job.Text)
		} else {
			fmt.Printf("%d. ○ %s\n", job.ID, job.Text)
		}
	}
}

func doneTask(list []Todo, targetID int) []Todo {
	for i := range list {
		if list[i].ID == targetID {
			list[i].Done = true
			break
		}
	}
	return list
}

func deleteTask(list []Todo, targetID int) []Todo {
	for i := range list {
		if list[i].ID == targetID {
			return append(list[:i], list[i+1:]...)
		}
	}
	return list
}
