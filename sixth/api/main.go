package main

import (
	"api/models"
	r "api/routes"
	"fmt"
)

func main() {
	fmt.Println("Iniciando o servidor Rest com Go")
	models.Accounts = []models.Account{
		{ID: "1", Name: "John Doe", Email: "john.doe@mail.com", Password: "123456"},
		{ID: "2", Name: "Bob Mate", Email: "bob.mate@mail.com", Password: "654321"},
		{ID: "3", Name: "Jennifer Lo", Email: "j.lo@mail.com", Password: "753951"},
	}
	r.HandleRequest()
}
