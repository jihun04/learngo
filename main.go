package main

import (
	"fmt"
	"learngo/accounts"
)

func checkErr(err error) {
	if err != nil {
		fmt.Println(err)
	}
}

func main() {
	account := accounts.NewAccount("jihun04")
	account.ChangeOwner("nine")
	fmt.Println(account)
}
