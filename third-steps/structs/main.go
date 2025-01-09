package structs

import "fmt"

type Account struct {
	ID       int
	Name     string
	Email    string
	Password string
}

func (account *Account) AccountLogin(password string) bool {
	return password == account.Password
}

func ShowAccountStruct() {
	john := Account{ID: 1, Name: "John Doe", Email: "john.doe@email.com", Password: "123456"}
	bob := Account{2, "Bob Smith", "bob.smith@email.com", "123456"}
	println(&john)
	println(&bob)
	johnStr := fmt.Sprintf("%+v", john)
	bobStr := fmt.Sprintf("%+v", bob)
	println(johnStr)
	println(bobStr)
}

func ShowAccounStructWithPointers() {
	kilian := new(Account)
	kilian.ID = 3
	kilian.Name = "Kilian"
	kilian.Email = "kilian1@email.com"
	kilian.Password = "123456"
	fmt.Println(*kilian)
}

func CompareStructs() {
	john := Account{ID: 1, Name: "John Doe", Email: "", Password: "123456"}
	bob := Account{ID: 1, Name: "John Doe", Email: "", Password: "123456"}

	// Here the comparison is about the content in the struct (properties)
	println("Comparing content")
	if john == bob {
		println("John and Bob are the same")
	} else {
		println("John and Bob are different")
	}
	println("\n\n\nComparing memory address")

	var kilian *Account
	kilian = new(Account)
	kilian.ID = 3
	kilian.Name = "Kilian"
	kilian.Email = "kilian@mail.com"
	kilian.Password = "123456"

	var kilian2 *Account
	kilian2 = new(Account)
	kilian2.ID = 3
	kilian2.Name = "Kilian"
	kilian2.Email = "kilian@mail.com"
	kilian2.Password = "123456"

	// Here the comparison is about the memory address of the struct
	if kilian == kilian2 {
		println("Kilian and Kilian2 are the same", "\nAddress of Kilian: ", &kilian, "\nAddress of Kilian2: ", &kilian2, "\n")
		fmt.Println("Content of kilian", *kilian, "\nContent of kilian2", *kilian2)
	} else {
		println("Kilian and Kilian2 are different", "\nAddress of Kilian: ", &kilian, "\nAddress of Kilian2: ", &kilian2, "\n")
		fmt.Println("Content of kilian", *kilian, "\nContent of kilian2", *kilian2)
	}
}

func Login(account *Account, password string) bool {
	return account.AccountLogin(password)
}
