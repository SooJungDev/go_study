package main

import (
	"fmt"
	"goForBeginner/accounts"
	"log"
)

func main() {
	account := accounts.NewAccount("soojung")
	account.Deposit(10)
	fmt.Println(account.Balance())
	err := account.Withdraw(20)
	if err != nil {
		log.Println(err)
	}
	fmt.Println(account)
}
