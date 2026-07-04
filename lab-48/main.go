package main

import "fmt"

type BankAccount struct {
	Owner   string
	Balance float64
}

func main() {
	account := BankAccount{Owner: "Ali", Balance: 0.00}
	ownerName, accBalance := account.PrintStatement()
	fmt.Printf("Account: %s\n", ownerName)
	fmt.Printf("Balance: PKR %.2f\n\n", accBalance)

	deposit := account.Deposit(5000.00)
	fmt.Printf("Deposited PKR %.2f\n", 5000.00)
	fmt.Printf("Balance: PKR %.2f\n\n", deposit)

	withdrawal1, _ := account.Withdraw(2000.00)
	fmt.Printf("Withdrew PKR %.2f\n", 2000.00)
	fmt.Printf("Balance: PKR %.2f\n\n", withdrawal1)

	withdrawl2, err := account.Withdraw(5000.00)
	if err != "" {
		fmt.Println(err)
		fmt.Printf("Balance: PKR %.2f\n\n", withdrawl2)
	} else {
		fmt.Printf("Withdrew PKR %.2f\n", 5000.00)
		fmt.Printf("Balance: PKR %.2f\n\n", withdrawl2)
	}
}

func (acc *BankAccount) Deposit(amountDepo float64) (amount float64) {
	acc.Balance += amountDepo
	return acc.Balance
}

func (acc *BankAccount) Withdraw(amountWdr float64) (amount float64, errMsg string) {
	if amountWdr >= acc.Balance {
		errMsg = fmt.Sprintf("Error: insufficient funds (tried to withdraw PKR %.2f, balance is PKR %.2f)", amountWdr, acc.Balance)
		return acc.Balance, errMsg
	}
	acc.Balance -= amountWdr
	return acc.Balance, ""
}

func (acc BankAccount) PrintStatement() (owner string, balance float64) {
	return acc.Owner, acc.Balance
}
