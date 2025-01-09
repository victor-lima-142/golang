package main

import (
	"fmt"
	"third-steps/structs"
)

func main() {
	myAccount := new(structs.Account)
	myAccount.ID = 1
	myAccount.Name = "John Doe"
	myAccount.Email = "john@mail.com"
	myAccount.Password = "123456"
	logged := myAccount.AccountLogin("123456")
	fmt.Println("Logged in:", logged)
}
