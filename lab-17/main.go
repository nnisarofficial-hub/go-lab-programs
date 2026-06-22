package main

import "fmt"

func main() {
	var (
		username = "admin"
		password = "go123"
	)

	var inputUser, inputPasswd string
	fmt.Print("Username: ")
	fmt.Scan(&inputUser)
	fmt.Print("Password: ")
	fmt.Scan(&inputPasswd)

	if inputUser == username && inputPasswd == password {
		fmt.Print("Login successful!")
	} else {
		fmt.Print("Invalid credentials")
	}
}
