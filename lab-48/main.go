package main

import (
	"errors"
	"fmt"
)

func main() {
	account := BankAccount{Owner: "Ali", Balance: 0.00}
	account.PrintStatement()
	deposit, err := account.Deposit(5000.00)
	if err != nil {
		fmt.Println("Error:", err)
		fmt.Printf("Balance: PKR %.2f\n\n", deposit)
	} else {
		fmt.Printf("Deposited PKR %.2f\n", 5000.00) 
		fmt.Printf("Balance: PKR %.2f\n\n", deposit)
	}

	withdrawal1, _ := account.Withdraw(2000.00)
	fmt.Printf("Withdrew PKR %.2f\n", 2000.00)
	fmt.Printf("Balance: PKR %.2f\n\n", withdrawal1)

	withdrawl2, err := account.Withdraw(5000.00)
	if err != nil {
		fmt.Println(err)
		fmt.Printf("Balance: PKR %.2f\n\n", withdrawl2)
	} else {
		fmt.Printf("Withdrew PKR %.2f\n", 5000.00)
		fmt.Printf("Balance: PKR %.2f\n\n", withdrawl2)
	}
}

type BankAccount struct {
	Owner   string
	Balance float64
}

func (acc *BankAccount) Deposit(amountDepo float64) (float64, error) {
	if amountDepo <= 0 {
		return acc.Balance, errors.New("deposit amount cannot be negative")
	}
	acc.Balance += amountDepo
	return acc.Balance, nil
}

func (acc *BankAccount) Withdraw(amountWdr float64) (float64, error) {
	if amountWdr <= 0 {
		return acc.Balance, errors.New("withdrawal amount must be greater than zero")
	}
	if amountWdr > acc.Balance {
		err := fmt.Errorf("error: insufficient funds (tried to withdraw PKR %.2f, balance is PKR %.2f)", amountWdr, acc.Balance)
		return acc.Balance, err
	}
	acc.Balance -= amountWdr
	return acc.Balance, nil
}

func (acc BankAccount) PrintStatement() {
	fmt.Printf("Account: %s\nBalance: %.2f\n\n", acc.Owner, acc.Balance)
}
