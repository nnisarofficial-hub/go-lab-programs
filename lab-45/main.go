package main

import (
	"bufio"
	"errors"
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
	store := map[string]string{}
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
	name = strings.ToLower(strings.TrimSpace(name))
	caser := cases.Title(language.English)
	displayName := caser.String(name)
	number, found := store[name]
	if found {
		fmt.Printf("%s: %s\n", displayName, number)
	} else {
		fmt.Printf("Contact '%s' not found.\n", displayName)
	}
}

func handleAdd(reader *bufio.Reader, store map[string]string) {
	fmt.Print("Enter name to add: ")
	name, err := reader.ReadString('\n')
	if err != nil {
		fmt.Print(err)
		return
	}
	name = strings.ToLower(strings.TrimSpace(name))
	fmt.Print("Enter number: ")
	number, err := reader.ReadString('\n')
	if err != nil {
		fmt.Print(err)
		return
	}
	number = strings.TrimSpace(number)
	err = addContact(store, name, number)
	if err != nil {
		fmt.Println(err)
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
	name = strings.ToLower(strings.TrimSpace(name))
	err = deleteContact(store, name)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("Contact name %s deleted successfully!\n", name)
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

func addContact(store map[string]string, name, number string) error {
	_, exists := store[name]
	if exists {
		return errors.New("error: This contact name already exists!")
	}
	store[name] = number
	return nil
}

func deleteContact(store map[string]string, name string) error {
	_, exists := store[name]
	if exists {
		delete(store, name)
		return nil
	}
	return errors.New("error: This contact not exists")
}
