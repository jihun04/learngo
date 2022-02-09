package accounts

import (
	"errors"
	"log"
)

// Account struct
type Account struct {
	owner   string
	balance int
}

// NewAccount creates Account
func NewAccount(owner string) *Account {
	account := Account{owner: owner, balance: 0}
	return &account
}

// Deposit x amount on your account
func (a *Account) Deposit(amount int) error {
	if amount < 0 {
		return errors.New("can't deposit. amount must be more than 0")
	}
	a.balance += amount
	return nil
}

// Return balance
func (a *Account) Balance() int {
	return a.balance
}

// Return owner
func (a *Account) Owner() string {
	return a.owner
}

// Withdraw balance
func (a *Account) Withdraw(amount int) error {
	if amount < 0 {
		return errors.New("can't withdraw. amount must be more than 0")
	}
	if amount >= a.balance {
		return errors.New("can't. you're too poor to withdraw as that amount")
	}
	a.balance -= amount
	return nil
}

// Change owner
func (a *Account) ChangeOwner(newOwner string) {
	a.owner = newOwner
}

// Called by go
func (a *Account) String() string {
	log.Fatalln(errors.New("err"))
	return "Whatever"
}
