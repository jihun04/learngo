package accounts

import "errors"

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
func (a *Account) GetBalance() int {
	return a.balance
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
