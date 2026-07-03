package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

func userMenu() {
	fmt.Println("Phone Book")
	fmt.Println("1. Look up")
	fmt.Println("2. Add")
	fmt.Println("3. Delete")
	fmt.Println("4. List all")
	fmt.Println("5. Quit")
}
func main() {
	userMenu()
	reader := bufio.NewReader(os.Stdin)
	store := dataStore()
	for {
		fmt.Print("> ")
		userChoice, err := reader.ReadString('\n')
		if err != nil {
			fmt.Print(err)
			return
		}
		choice := strings.TrimSpace(userChoice)
		switch choice {
		case "1":
			handleLookUp(reader, store)
		case "2":
			handleAdd(reader, store)
		case "3":
			handleDelete(reader, store)
		case "4":
			handleListAll(store)
		case "5":
			return
		default:
			fmt.Println("Invalid choice, please try again.")
		}
	}
}

// =======================
// Handling Core functions
// =======================

func handleLookUp(reader *bufio.Reader, store map[string]string) {
	fmt.Print("Enter name: ")
	name, err := reader.ReadString('\n')
	if err != nil {
		fmt.Print(err)
		return
	}
	nameClean := strings.ToLower(strings.TrimSpace(name))
	caser := cases.Title(language.English)
	displayName := caser.String(nameClean)
	number := store[nameClean]
	if number == "" {
		fmt.Printf("Contact '%s' not found.\n", displayName)
	} else {
		fmt.Printf("%s: %s\n", displayName, number)
	}
}

func handleAdd(reader *bufio.Reader, store map[string]string) {
	fmt.Print("Enter name to add: ")
	name, err := reader.ReadString('\n')
	if err != nil {
		fmt.Print(err)
		return
	}
	nameClean := strings.ToLower(strings.TrimSpace(name))
	fmt.Print("Enter number: ")
	number, err := reader.ReadString('\n')
	if err != nil {
		fmt.Print(err)
		return
	}
	numberClean := strings.TrimSpace(number)
	errMsg := addContact(store, nameClean, numberClean)
	if errMsg != "" {
		fmt.Println(errMsg)
	} else {
		fmt.Print("Number added successfully!\n")
	}
}

func handleDelete(reader *bufio.Reader, store map[string]string) {
	fmt.Print("Enter name to delete: ")
	name, err := reader.ReadString('\n')
	if err != nil {
		fmt.Print(err)
		return
	}
	nameClean := strings.ToLower(strings.TrimSpace(name))
	errMsg := deleteContact(store, nameClean)
	if errMsg != "" {
		fmt.Println(errMsg)
	} else {
		fmt.Printf("Contact name %s deleted successfully!\n", nameClean)
	}
}

func handleListAll(store map[string]string) {
	caser := cases.Title(language.English)
	for name, number := range store {
		capitalName := caser.String(name)
		fmt.Printf("%s: %s\n", capitalName, number)
	}
}

// ===================
// Core Data Function
// ===================

func dataStore() map[string]string {
	contactDetails := map[string]string{
		"ali":  "0300-1234567",
		"sara": "0321-9876543",
		"umar": "0333-1111222",
	}
	return contactDetails
}

func addContact(store map[string]string, name, number string) string {
	_, exists := store[name]
	if exists {
		return "Error: This contact name already exists!"
	}
	store[name] = number
	return ""
}

func deleteContact(store map[string]string, name string) string {
	_, exists := store[name]
	if exists {
		delete(store, name)
		return ""
	}
	return "Error: This contact not exists"
}
