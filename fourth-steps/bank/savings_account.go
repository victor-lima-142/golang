package bank

import (
	"errors"
	"fourth-steps/employees"
)

type SavingsAccount struct {
	Employee      *employees.Employee
	balance       int
	accountNumber int
}

func (s *SavingsAccount) GetBalance() int {
	return s.balance
}

func (s *SavingsAccount) SetBalance(balance int) {
	s.balance = balance
}

func (s *SavingsAccount) Withdraw(amount int) {
	newValue := s.balance - amount
	s.SetBalance(newValue)
}

func (s *SavingsAccount) SetAccountNumber(accountNumber int) {
	s.accountNumber = accountNumber
}

func (s *SavingsAccount) GetAccountNumber() int {
	return s.accountNumber
}

func (s *SavingsAccount) Deposit(amount int, destination Account) (bool, error) {
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
