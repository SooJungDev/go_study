package accounts

import (
	"errors"
	"fmt"
)

// Account Struct
type Account struct {
	owner   string
	balance int
}

var errNoMoney = errors.New("Can't withdraw")

// NewAccount creates Account
func NewAccount(owner string) *Account {
	account := Account{owner: owner, balance: 0}
	return &account
}

// 복사본을 만들지 않고 balance 를 변경해야함
// 누군가 호출 한다면 account를 복사해서 사용하는 것이 아니라 호출한 account 를 사용 하라는 뜻
// Deposit x amount on your account
func (a *Account) Deposit(amount int) {
	a.balance += amount
}

// Balance of Your account
func (a Account) Balance() int {
	return a.balance
}

// Withdraw x amount from your account
func (a *Account) Withdraw(amount int) error {
	if a.balance < amount {
		return errNoMoney
	}
	a.balance -= amount
	return nil
}

// ChangeOwner of the account
func (a *Account) ChangeOwner(newOwners string) {
	a.owner = newOwners
}

// Owner of the account
func (a Account) Owner() string {
	return a.owner
}

func (a Account) String() string {
	return fmt.Sprint(a.Owner(), "'s account.\nHas: ", a.balance)
}
