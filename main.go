package main

import (
	"fmt"
	"learngo/accounts"
	"log"
)

func checkErr(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}

func main() {
	account := accounts.NewAccount("jihun")
	checkErr(account.Deposit(10))
	checkErr(account.Withdraw(11))
	fmt.Println(account.GetBalance())
}
