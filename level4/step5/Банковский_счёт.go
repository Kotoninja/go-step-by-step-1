package main

import (
	"errors"
	"fmt"
)

type Account struct {
	balance float64
	owner   string
}

func NewAccount(owner string) *Account {
	return &Account{owner: owner, balance: 0}
}

func (a *Account) SetBalance(value float64) error {
	if value < 0 {
		return errors.New("Value < 0")
	} else if a.balance < 0 {
		return errors.New("Balance < 0")
	}
	a.balance = value
	return nil
}

func (a Account) GetBalance() float64 {
	return a.balance
}

func (a *Account) Deposit(value float64) error {
	if value < 0 {
		return errors.New("Value < 0")
	} else if a.balance < 0 {
		return errors.New("Balance < 0")
	}

	a.balance += value

	return nil
}

func (a *Account) Withdraw(value float64) error {
	if value < 0 {
		return errors.New("Value < 0")
	} else if a.balance < 0 {
		return errors.New("Balance < 0")
	} else if a.balance-value < 0 {
		return errors.New("Balance - value < 0")
	}

	a.balance -= value

	return nil
}

func main() {
	account := NewAccount("Alice")

	err := account.SetBalance(1000.0)
	if err != nil {
		fmt.Println(err)
		// t.Errorf("Ошибка при установке баланса: %v", err)
	}

	// Вносим деньги
	err = account.Deposit(500.0)
	if err != nil {
		fmt.Println(err)
		// t.Errorf("Ошибка при внесении денег: %v", err)
	}

	// Снимаем деньги
	err = account.Withdraw(200.0)
	if err != nil {
		fmt.Println(err)
		// t.Errorf("Ошибка при снятии денег: %v", err)
	}

	// Получаем текущий баланс
	balance := account.GetBalance()
	fmt.Println(balance)
}
