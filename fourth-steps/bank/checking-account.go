package bank

import (
	"errors"
	"fourth-steps/employees"
)

type CheckingAccount struct {
	Employee      *employees.Employee
	balance       int
	accountNumber int
}

func (s *CheckingAccount) GetBalance() int {
	return s.balance
}

func (s *CheckingAccount) SetBalance(balance int) {
	s.balance = balance
}

func (s *CheckingAccount) Withdraw(amount int) {
	newValue := s.balance - amount
	s.SetBalance(newValue)
}

func (s *CheckingAccount) SetAccountNumber(accountNumber int) {
	s.accountNumber = accountNumber
}

func (s *CheckingAccount) GetAccountNumber() int {
	return s.accountNumber
}

func (s *CheckingAccount) Deposit(amount int, destination Account) (bool, error) {
	var err error
	if amount < 0 {
		err = errors.New("amount cannot be negative")
		return false, err
	} else if s.balance < amount {
		err = errors.New("insufficient funds")
		return false, err
	}
	currentBalance := destination.GetBalance()
	s.SetBalance(s.GetBalance() - amount)
	destination.SetBalance(currentBalance + amount)
	return true, nil
}
