package bank

type Account interface {
	SetAccountNumber(accountNumber int)
	GetAccountNumber() int
	SetBalance(balance int)
	GetBalance() int
	Deposit(amount int) (bool, error)
	Withdraw(amount int)
}
